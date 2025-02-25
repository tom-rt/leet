package main

import (
	"sync"
)

func maxAreaDumb(height []int) int {
	var maxArea int
	var left int = 0
	var right int = len(height) - 1
	for left < right {
		h := height[left]
		if h >= height[right] {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	var maxArea int
	c := make(chan int, len(height))
	var wg sync.WaitGroup

	for i := 0; i < len(height); i++ {
		wg.Add(1)
		go calcMaxAreaDumb(height, i, &wg, c)
	}

	wg.Wait()
	close(c)

	for area := range c {
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func calcMaxAreaDumb(batch []int, i int, wg *sync.WaitGroup, c chan int) {
	defer func() {
		wg.Done()
	}()
	var maxArea int
	for j := i + 1; j < len(batch); j++ {
		h := batch[i]
		if h >= batch[j] {
			h = batch[j]
		}
		area := h * (j - i)
		if area > maxArea {
			maxArea = area
		}
	}
	c <- maxArea
}
