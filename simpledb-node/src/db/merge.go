package db

import (
	"sort"
)

func (level *Level) mergeSort(files []string) ([]*LSMDataEntry, error) {
	if len(files) > 1 {
		mid := len(files) / 2

		left, err := level.mergeSort(files[:mid])
		if err != nil {
			return nil, err
		}
		right, err := level.mergeSort(files[mid:])
		if err != nil {
			return nil, err
		}

		result := mergeHelper(left, right)

		return result, nil
	}

	return level.fm.MMap(files[0])
}

func mergeHelper(left, right []*LSMDataEntry) (entries []*LSMDataEntry) {
	od := newOrderedDict()
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		leftEntry := left[i]
		rightEntry := right[j]

		if val, ok := od.Get(leftEntry.key); ok {
			if val.(*LSMDataEntry).seqID < leftEntry.seqID {
				od.Set(leftEntry.key, leftEntry)
			}
			i++
			continue
		}

		if val, ok := od.Get(rightEntry.key); ok {
			if val.(*LSMDataEntry).seqID < rightEntry.seqID {
				od.Set(rightEntry.key, rightEntry)
			}
			j++
			continue
		}

		if leftEntry.key < rightEntry.key {
			od.Set(leftEntry.key, leftEntry)
			i++
		} else if leftEntry.key == rightEntry.key {
			if leftEntry.seqID > rightEntry.seqID {
				od.Set(leftEntry.key, leftEntry)
			} else {
				od.Set(rightEntry.key, rightEntry)
			}
			i++
			j++
		} else {
			od.Set(rightEntry.key, rightEntry)
			j++
		}
	}

	for i < len(left) {
		leftEntry := left[i]
		if val, ok := od.Get(leftEntry.key); ok {
			if val.(*LSMDataEntry).seqID < leftEntry.seqID {
				od.Set(leftEntry.key, leftEntry)
			}
		} else {
			od.Set(leftEntry.key, leftEntry)
		}
		i++
	}

	for j < len(right) {
		rightEntry := right[j]
		if val, ok := od.Get(rightEntry.key); ok {
			if val.(*LSMDataEntry).seqID < rightEntry.seqID {
				od.Set(rightEntry.key, rightEntry)
			}
		} else {
			od.Set(rightEntry.key, rightEntry)
		}
		j++
	}

	for val := range od.Iterate() {
		entries = append(entries, val.(*LSMDataEntry))
	}

	return entries
}

func mergeIntervals(intervals []*merge) []*merge {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].keyRange.startKey < intervals[j].keyRange.startKey
	})

	result := []*merge{}

	for i, interval := range intervals {
		if i == 0 {
			result = append(result, interval)
		} else {
			curr := result[len(result)-1]
			currEnd := curr.keyRange.endKey
			nextStart := interval.keyRange.startKey
			nextEnd := interval.keyRange.endKey

			if nextStart > currEnd {
				result = append(result, interval)
			} else {
				if nextEnd > currEnd {
					curr.keyRange.endKey = nextEnd
				}
				curr.files = append(curr.files, interval.files...)
			}
		}
	}

	return result
}

// mergeInterval joins the interval into the list if it overlaps with an interval. Else it returns the original list.
func mergeInterval(intervals []*merge, newInterval *merge) (merged bool, result []*merge) {
	merged = false
	newLeft := newInterval.keyRange.startKey
	newRight := newInterval.keyRange.endKey
	newFiles := newInterval.files

	for i, interval := range intervals {
		left := interval.keyRange.startKey
		right := interval.keyRange.endKey

		if right < newLeft {
			result = append(result, interval)
		} else if (left <= newLeft && newLeft <= right) || (left <= newRight && newRight <= right) || (newLeft <= left && newRight >= right) {
			if left < newLeft {
				newLeft = left
			}
			if right > newRight {
				newRight = right
			}
			newFiles = append(newFiles, interval.files...)
			merged = true
		} else {
			if merged {
				result = append(result, &merge{
					files: newFiles,
					keyRange: &KeyRange{
						startKey: newLeft,
						endKey:   newRight,
					},
				})
				return true, append(result, intervals[i:]...)
			}
			return false, intervals
		}
	}

	if merged {
		return true, append(result, &merge{
			files: newFiles,
			keyRange: &KeyRange{
				startKey: newLeft,
				endKey:   newRight,
			},
		})
	}
	return false, intervals
}
