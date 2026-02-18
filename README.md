# haar-training

Go application to automate OpenCV Haar Cascade classifier training.

The program orchestrates the full pipeline: downloading negative images from ImageNet, generating positive samples, and training the classifier to produce an XML file usable for object detection.

## Prerequisites

- **Go** (GOPATH mode, no modules)
- **OpenCV** CLI tools:
  - `opencv_createsamples`
  - `opencv_traincascade`

These tools must be available in your `PATH`.

## Installation

```bash
go get github.com/marianina8/haar-training
cd $GOPATH/src/github.com/marianina8/haar-training
```

### Go Dependency

- [disintegration/imaging](https://github.com/disintegration/imaging) — image resizing and grayscale conversion

```bash
go get github.com/disintegration/imaging
```

## Usage

```bash
go run main.go
```

The program executes the following steps in order:

### 1. Download Negative Images

Fetches images from ImageNet (categories: plants, landscapes, crowds, buildings, etc.), resizes them to 200x200 pixels, converts them to grayscale, and stores them in `negatives/`.

### 2. Generate bg.txt

Lists all negative image paths in `bg.txt`, a file required by OpenCV for training.

### 3. Create Positive Samples

Calls `opencv_createsamples` to generate variations of the positive image (`adidas.png`) with rotations and distortions, producing the `positives.vec` file.

### 4. Train the Classifier

Calls `opencv_traincascade` to train the Haar Cascade model from positive and negative samples. The resulting XML classifier is written to the `data/` directory.

## Project Structure

```
├── main.go              # Entry point, pipeline orchestration
├── images/
│   └── images.go        # Image downloading and processing from ImageNet
├── samples/
│   ├── create_samples.go      # Positive sample generation (opencv_createsamples)
│   └── create_posvector.go    # Positive vector file creation
├── training/
│   └── cascade.go       # Haar Cascade training (opencv_traincascade)
├── adidas.png           # Example positive image
├── bg.txt               # Negative image list (generated)
├── negatives/           # Downloaded negative images (generated)
├── positives.vec        # Positive sample vector (generated)
└── data/                # Trained classifier output (generated)
```

## Training Parameters

Default parameters in `main.go`:

| Parameter | Value |
|---|---|
| Sample size | 70x70 px |
| Positive samples | 1950 (creation) / 1800 (training) |
| Negative samples | 900 |
| Number of stages | 20 |
| Negative image limit | 4000 |
| Negative image size | 200x200 px |

To train on a different object, replace `adidas.png` with your positive image and adjust the parameters in `main.go`.
