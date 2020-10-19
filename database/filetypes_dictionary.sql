TRUNCATE filetypes CASCADE;
SELECT setval('filetypes_type_id_seq', 1);

INSERT INTO filetypes 
VALUES(1, 'unknown', 'No file type was found in databse. Contact administrator for more info.');

INSERT INTO filetypes(type_extension, type_description) 
VALUES('.bmp', 'BMP or Bitmap Image File is a format developed by Microsoft for Windows. There is no compression or information loss with BMP files which allow images to have very high quality, but also very large file sizes. Due to BMP being a proprietary format, it is generally recommended to use TIFF files.');

INSERT INTO filetypes(type_extension, type_description) 
VALUES('.tiff', 'TIFF or Tagged Image File Format are lossless images files meaning that they do not need to compress or lose any image quality or information (although there are options for compression), allowing for very high-quality images but also larger file sizes.');
INSERT INTO filetypes(type_extension, type_description) 
VALUES('.tif', 'TIFF or Tagged Image File Format are lossless images files meaning that they do not need to compress or lose any image quality or information (although there are options for compression), allowing for very high-quality images but also larger file sizes.');

INSERT INTO filetypes(type_extension, type_description) 
VALUES('.jpeg', 'JPEG, which stands for Joint Photographic Experts Groups is a “lossy” format meaning that the image is compressed to make a smaller file. The compression does create a loss in quality but this loss is generally not noticeable. JPEG files are very common on the Internet and JPEG is a popular format for digital cameras - making it ideal for web use and non-professional prints.');
INSERT INTO filetypes(type_extension, type_description) 
VALUES('.jpg', 'JPEG, which stands for Joint Photographic Experts Groups is a “lossy” format meaning that the image is compressed to make a smaller file. The compression does create a loss in quality but this loss is generally not noticeable. JPEG files are very common on the Internet and JPEG is a popular format for digital cameras - making it ideal for web use and non-professional prints.');

INSERT INTO filetypes(type_extension, type_description) 
VALUES('.png', 'PNG or Portable Network Graphics files are a lossless image format originally designed to improve upon and replace the gif format. PNG files are able to handle up to 16 million colors, unlike the 256 colors supported by GIF.');

INSERT INTO filetypes(type_extension, type_description) 
VALUES('.gif', 'GIF or Graphics Interchange Format files are widely used for web graphics, because they are limited to only 256 colors, can allow for transparency, and can be animated. GIF files are typically small is size and are very portable.');

SELECT * FROM filetypes;