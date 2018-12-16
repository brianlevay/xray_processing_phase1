# X-Ray Image Processing for Cores at IODP

This software is intended to allow scientists to process x-ray images acquired from cores. It's designed around the initial setup used on the JOIDES resolution during Phase 1 of the X-Ray Imager Project.

## Critical Assumptions That Must Be Met
* The core is expected to be oriented roughly top to bottom in the image. For the detector that's currently being used in Phase 1, that corresponds to the core being oriented along the long axis of the detector. The image searching algorithm uses a row-first approach, and it will fail if the core is aligned left to right. Future efforts may allow this to be configurable. 
* The source is assumed to be centered over the detector for all projection calculations. This doesn't need to be exactly true in real life, but the more off-center the source becomes, the more error will be introduced. Future efforts may allow this to be set as part of the configuration.
* The core is assumed to be roughly cylindrical. The code will not work correctly on other types of shapes, such as slabs.

## Basic User Workflow
1. Select the raw images to process
2. Generate a combined histogram for pixel values in the images
3. Select the desired histogram bounds to be used for processing
4. Specify the geometry of the system and the core
5. Process the images

## Basic Processing Workflow (Performed by the Software)
1. Open the raw image
2. Detect the location of the core in the image (unless the location is specified by the user)
3. Model the x-ray path from the source to each pixel using the specified location and geometry
4. Calculate &mu;&rho;t from the raw image (primary image enhancement step)
5. Perform thickness compensation using modelled ray paths to "flatten" the core
6. Apply contrast enhancement algorithms to improve the image
7. Add scale bars
8. Save the new image

## Important Processing Notes
* Grayscale values are globally consistent; a pixel value in one part of an image has the same underlying meaning as the same value in another part of the image
* The relationship between the input grayscale and the output grayscale is monotonic within the bounds specified by the user
* The relationship between the input grayscale and the output grayscale is the same for all images processed in the same batch

## Installing and Running the Software
*coming soon*