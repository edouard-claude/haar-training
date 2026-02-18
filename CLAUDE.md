# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go application that automates OpenCV Haar Cascade classifier training. It orchestrates a 3-step pipeline:
1. Download negative (background) images from ImageNet
2. Generate positive sample vectors using `opencv_createsamples`
3. Train a Haar Cascade classifier using `opencv_traincascade`

## Build & Run

This project uses legacy GOPATH-based Go (no go.mod). The import path is `github.com/marianina8/haar-training`.

```bash
# Must be in GOPATH at $GOPATH/src/github.com/marianina8/haar-training
go build
go run main.go
```

**External dependency:** `github.com/disintegration/imaging` (used in `images/` package for resize/grayscale).

**System requirements:** OpenCV command-line tools (`opencv_createsamples`, `opencv_traincascade`) must be on PATH.

## Architecture

- **`main.go`** — Pipeline orchestrator. Configures ImageNet URLs, downloads negatives, generates `bg.txt` (list of negative image paths), then calls sample creation and training.
- **`images/`** — `images.Get()` downloads images from ImageNet URLs, resizes them, optionally converts to grayscale, saves as numbered JPGs.
- **`samples/`** — Wrappers around `opencv_createsamples`. `CreateSamples()` generates distorted positive samples into `positives.vec`. `CreatePositiveVectorFile()` creates a vector file from an info list.
- **`training/`** — `HaarCascade()` wraps `opencv_traincascade`, streams stderr output during training.

## Key Generated Artifacts

- `negatives/` — Downloaded background images (numbered JPGs)
- `bg.txt` — Newline-separated list of negative image paths
- `positives.vec` — OpenCV sample vector file
- `data/` — Output directory for trained cascade XML
- `info/` — Intermediate sample info directory
