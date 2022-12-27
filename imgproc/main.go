package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {
	srcDir := flag.String("src", "", "Input dir")
	dstDir := flag.String("dst", "", "Output dir")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgrp", "waitgrp/channel")
	poolSize := flag.Int("poolsize", 4, "nb worker")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize )
	}

	start := time.Now()
	err := t.Process()
	if err != nil {
		fmt.Printf("an error happened: %v", err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Println("Processing took: ", elapsed)
}
