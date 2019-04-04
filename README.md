# X-Ray Image Processing for Cores at IODP

This software is intended to allow scientists to process x-ray images acquired from cores. It's designed around the initial setup used on the JOIDES resolution during Phase 1 of the X-Ray Imager Project.

## Installing and Running the Software
1. Compile the code for your target architecture (use the compile.sh script as a reference, if needed)
2. If you have admin rights to your computer, copy the executable and the "static" folder to the directory of your choice
3. If you don't have admin rights, copy the executable and "static" folder to the root of the directory where your raw images will be stored
4. Run the executable, open your web browser, and navigate to localhost at the port specified in the terminal display

## Library Dependencies Not in Repository
* golang.org/x/image/tiff