# Kairos-SDK-Go

[![Vexor status](https://ci.vexor.io/projects/96cfab2e-74e0-4e21-8812-2e05ac6047aa/status.svg)](https://ci.vexor.io/ui/projects/96cfab2e-74e0-4e21-8812-2e05ac6047aa/builds)

This is the Go wrapper for the [Kairos Facial Recognition API](https://www.kairos.com). The wrapper was written for [Humaniq.co](https://humaniq.co/) project needs, during the facial image research process. 

## Contributors
This  wrapper was developed by:

* [Alexander Kirillov](https://github.com/saratovsource)
* [Kareem Hepburn](https://github.com/magicalbanana) 



---


Ported methods:
====================

## Authenticate Once

Before you can make API calls you'll need to pass Kairos your credentials **App Id** and **App Key** (You only need to do this once). Paste your App Id and App Key into the constructor method like so:

```
// Alexander to fill
```


    
## View Your Galleries

This method returns a list of all galleries you've created:

```
// Alexander to fill
```

## View Your Subjects

This method returns a list of all subjects for a given gallery:

```
// Alexander to fill
```

## Remove a Subject

This method removes a subject from given gallery:

```
// Alexander to fill
```

## Remove a Gallery

This method removes a given gallery:

```
// Alexander to fill
```

## Enroll an Image

The **Enroll** method **registers a face for later recognitions**. Here's an example of enrolling a face (subject) using a method that accepts an image URL or image data in base64 format, and enrolls it as a new subject into your specified gallery:    

```
// Alexander to fill
```
`The SDK also includes a file upload field, which converts a local image file to base64 data.`

## Recognize an Image

The **Recognize** method takes an image of a subject and **attempts to match it against a given gallery of previously-enrolled subjects**. Here's an example of recognizing a subject using a method that accepts an image URL or image data in base64 format, sends it to the API, and returns a match and confidence value:    

```
// Alexander to fill
```

`The SDK also includes a file upload field, which converts a local image file to base64 data.`

## Detect Image Attributes

The **Detect** method takes an image of a subject and **returns various attributes pertaining to the face features**. Here's an example of using detect via method that accepts an image URL or image data in base64 format, sends it to the API, and returns face attributes:    

```
// Alexander to fill
```

`The SDK also includes a file upload field, which converts a local image file to base64 data.`

## Verify image

The **Verify** method takes an image and verifies that it matches an existing subject in a gallery.  Here's an example of using verify via method that accepts a path to an image file, sends it to the API, and returns face attributes: 

```
// Alexander to fill
```
`The SDK also includes a file upload field, which converts a local image file to base64 data.`



---


## Copyrights

All rights belong [Kairos support page](http://www.kairos.com/support). 
