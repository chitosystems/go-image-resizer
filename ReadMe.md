# Image Resizer Script (GoLang)

This GoLang script takes a URL as input, resizes the image, and returns it as a base64 JSON object with the following attributes:

- `image`: The resized image in base64 format.
- `message`: Informational messages or errors (if any).
- `status`: The status of the operation (success or failure).

## Usage

To use this script, provide a URL with the image you want to resize in the following format:


     http://localhost:3000/?img=<IMAGE_URL>&w=<WIDTH>&h=<HEIGHT>



Replace `<IMAGE_URL>` with the URL of the image you want to resize and `<WIDTH>` and `<HEIGHT>` with the desired width and height.

## Example

Suppose you want to resize an image from the following URL:

     http://localhost:3000/?img=https://domain.online/assets/d7e9a118af5806d5a.jpeg&w=300&h=300


## Usage

You can make a request to the script, and it will return a JSON response with the resized image:


```php
<?php

$ch = curl_init();

curl_setopt($ch, CURLOPT_URL, 'http://localhost:3000/?img=https://cdn.mafaro.online/assets/artists/36c25ec1d00ce0bd7e9a118af5806d5a.jpeg&w=300&h=300');
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

$result = curl_exec($ch);
if (curl_errno($ch)) {
    echo 'Error:' . curl_error($ch);
}
curl_close($ch);
```


```json
{
  "base64_image": "base64_encoded_image_data",
  "message": "Image resized successfully.",
  "status": "success"
}
```

```html
<!-- html usage -->
<img src="$base64_image" />
```

### Dependencies


[disintegration/imaging](https://github.com/disintegration/imaging)

This GoLang script may use external packages or libraries for HTTP handling and image resizing. Be sure to list and provide instructions for installing these dependencies.

### License
This script is licensed under the MIT License. 


In this version, I've assumed that your GoLang script follows a typical project structure with a `main.go` file that you build and run. Be sure to replace `<repository_url>` with the actual URL of your script repository and provide appropriate installation and configuration instructions specific to your GoLang project.
