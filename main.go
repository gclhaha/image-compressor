package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
)

var (
	targetSizeKB = flag.Int("s", 1000, "Target size in KB for JPEG compression (default 1000)")
	maxWidth     = flag.Int("w", 1920, "Maximum width in Px for image resizing (default 1920)")
	inputDir     string
	outputDir    string
)

func main() {
	flag.StringVar(&inputDir, "i", "", "Input directory (required)")
	flag.StringVar(&outputDir, "o", "", "Output directory (required)")
	flag.Parse()

	// Validate required flags
	if inputDir == "" || outputDir == "" {
		flag.PrintDefaults()
		return
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Traverse the input directory
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		// Check if the file is an image (jpg or png)
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			// Load the image file
			imgFile, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening %s: %v\n", path, err)
				return nil
			}
			defer imgFile.Close()

			// Decode the image
			img, format, err := image.Decode(imgFile)
			if err != nil {
				fmt.Printf("Error decoding %s: %v\n", path, err)
				return nil
			}

			// Resize the image if it exceeds maxWidth
			resizedImg := resizeImage(img, *maxWidth)

			// Determine the relative path for the output file
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				fmt.Printf("Error determining relative path for %s: %v\n", path, err)
				return nil
			}
			outPath := filepath.Join(outputDir, relPath)

			// Create output directory if it doesn't exist
			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				fmt.Printf("Error creating output directory %s: %v\n", filepath.Dir(outPath), err)
				return nil
			}

			// Create output file
			outFile, err := os.Create(outPath)
			if err != nil {
				fmt.Printf("Error creating output file %s: %v\n", outPath, err)
				return nil
			}
			defer outFile.Close()

			// Compress and save the resized image
			if format == "jpeg" || format == "jpg" {
				err = compressJPEGImage(resizedImg, outFile, *targetSizeKB)
			} else if format == "png" {
				err = compressPNGImage(resizedImg, outFile)
			}

			if err != nil {
				fmt.Printf("Error compressing %s: %v\n", path, err)
				return nil
			}

			fmt.Printf("Compressed %s successfully\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", inputDir, err)
		return
	}
}

func compressJPEGImage(img image.Image, outFile *os.File, targetSizeKB int) error {
	var buf bytes.Buffer
	quality := 80

	for quality > 0 {
		buf.Reset()
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return err
		}

		if buf.Len() <= targetSizeKB*1024 {
			break
		}

		quality -= 5
	}

	_, err := outFile.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func compressPNGImage(img image.Image, outFile *os.File) error {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return err
	}

	_, err = outFile.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func resizeImage(img image.Image, maxWidth int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width <= maxWidth {
		return img
	}

	// Calculate proportional height
	newWidth := maxWidth
	newHeight := (newWidth * height) / width

	// Resize image using Lanczos resampling
	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, bounds, draw.Over, nil)

	return resizedImg
}