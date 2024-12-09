package day09

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Day09 struct{}

var _ solver.Solver = (*Day09)(nil)

func New() *Day09 {
	return &Day09{}
}

type File struct {
	ID, StartIdx, EndIdx, Length int
}

func NewFile(id, startIdx, endIdx, length int) File {
	return File{id, startIdx, endIdx, length}
}

func (d *Day09) Part1(lines []string) string {
	disk, _ := d.parseDisk(lines)

	d.compact(disk)

	checkSum := d.getCheckSum(disk)

	return fmt.Sprintf("%d", checkSum)
}

func (d *Day09) Part2(lines []string) string {
	disk, files := d.parseDisk(lines)

	disk = d.compactWholeFile(disk, files)

	checkSum := d.getCheckSum(disk)

	return fmt.Sprintf("%d", checkSum)
}

func (d *Day09) parseDisk(lines []string) ([]int, []File) {
	diskMap, _ := util.ReadInts(lines[0], "")

	id := 0
	disk := make([]int, 0)
	files := make([]File, 0) // fileID is the idx and the value is the size

	startIDIdx := 0
	for idx, fileLength := range diskMap {
		if idx%2 == 0 {
			var freeSpace int
			if idx != len(diskMap)-1 {
				freeSpace = diskMap[idx+1]
			}
			for i := 0; i < fileLength; i++ {
				disk = append(disk, id)
			}
			for i := 0; i < freeSpace; i++ {
				disk = append(disk, -1) // -1 will represent free memory
			}
			// files represents a file
			// startIDIDx is the start of the ID such that we can use it to do the swaps easier
			files = append(files, NewFile(id, startIDIdx, startIDIdx+fileLength, fileLength))
			id++

			startIDIdx += fileLength + freeSpace
		}
	}
	return disk, files
}

func (d *Day09) compact(disk []int) {
	// start from end of disk
	for i := len(disk) - 1; i >= 0; i-- {
		// free memory, skip it
		if disk[i] == -1 {
			continue
		}

		// find first free space from left
		for j := 0; j < i; j++ {
			if disk[j] == -1 {
				//swap
				disk[j] = disk[i]
				disk[i] = -1
				break
			}
		}
	}
}

func (d *Day09) compactWholeFile(disk []int, files []File) []int {
	// start from last file ID
	for fileID := len(files) - 1; fileID >= 0; fileID-- {
		if fileID == 0 {
			break
		}
		// the file size we're trying to find space for
		file := files[fileID]

		// now we need to find first chunk of free memory
		freeMemoryStart := -1
		freeMemoryLength := 0

		// start from beginning of disk
		for i := 0; i < len(disk); i++ {
			// we found free space!
			if disk[i] == -1 {
				if freeMemoryStart == -1 {
					freeMemoryStart = i
				}

				// stop computing if we're overlapping
				if freeMemoryStart > file.StartIdx {
					break
				}

				freeMemoryLength++
				if freeMemoryLength == file.Length { // file size fits within memory chunk and we haven't start to overlap
					fileIDsToSwap := disk[file.StartIdx:file.EndIdx]
					freeMemory := disk[freeMemoryStart : freeMemoryStart+freeMemoryLength]

					// copy a new disk
					// [0, 0, -1, -1, 9, 9 -1]
					newDisk := make([]int, 0)
					newDisk = append(newDisk, disk...)

					// copy over from the start of the disk to where the free memory is
					// newDisk = [0, 0, -1, -1, 9, 9, -1]
					// [0, 0] - newDisk[:freeMemoryStart]- copying all elements until the free memory
					// [0, 0, 9, 9] - append(fileIDsToSwap, ...) - append the fileIDs we're swapping as well as all other elements after it
					// [0, 0, 9, 9, 9, 9, -1] - newDisk[freeMemoryStart+freeMemoryLength:]... - everything else afterwards
					// then we'll append the fileIDs that we're swapping then append the rest of the disk
					newDisk = append(newDisk[:freeMemoryStart], append(fileIDsToSwap, newDisk[freeMemoryStart+freeMemoryLength:]...)...)

					// same thing but now we're swapping the free memory
					// newDisk = [0, 0, 9, 9, 9, 9, -1]
					// [0, 0, 9, 9] - newDisk[:file.StartIdx] - copying all elements until the the start ID IDX
					// [0, 0, 9, 9, -1, -1] - append(freeMemory, ...) - append the free memory we're swapping
					// [0, 0, 9, 9, 9, 9, -1] - newDisk[freeMemoryStart+freeMemoryLength:]... - everything else afterwards
					newDisk = append(newDisk[:file.StartIdx], append(freeMemory, newDisk[file.EndIdx:]...)...)

					// overwrite disk
					disk = newDisk
					break
				}
			} else {
				// reset since we did not find enough free space
				freeMemoryStart = -1
				freeMemoryLength = 0
			}
		}
	}
	return disk
}

func (d *Day09) getCheckSum(disk []int) int {
	checkSum := 0
	for blockPos, fileID := range disk {
		if fileID != -1 {
			checkSum += blockPos * fileID
		}
	}
	return checkSum
}
