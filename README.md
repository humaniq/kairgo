# Kairos-SDK-Go

[![Vexor status](https://ci.vexor.io/projects/96cfab2e-74e0-4e21-8812-2e05ac6047aa/status.svg)](https://ci.vexor.io/ui/projects/96cfab2e-74e0-4e21-8812-2e05ac6047aa/builds)

This is the Go wrapper for the [Kairos Facial Recognition API](https://www.kairos.com). The wrapper was written for [Humaniq.co](https://humaniq.co/) project needs, during the facial image research process. 

## Contributors
This  wrapper was developed by:

* [Alexander Kirillov](https://github.com/saratovsource)
* [Kareem Hepburn](https://github.com/magicalbanana) 



---

Usage
=====

```go
import "github.com/humaniq/kairgo"
```

Create a new Kairos client, then use the exposed services to access
different parts of the Kairos API.

Ported methods:
====================

## Authenticate Once

Before you can make API calls you'll need to pass Kairos your credentials **App Id** and **App Key** (You only need to do this once). Paste your App Id and App Key into the constructor method like so:

```go
import "github.com/humaniq/kairgo"

const (
  APP_ID  = "kairos_app_id"
  APP_KEY = "kairos_app_key"
)

client, err := kairgo.New("", APP_ID, APP_KEY, nil)
```


    
## View Your Galleries

This method returns a list of all galleries you've created:

```go
gallery, err := client.ViewGallery("MyGallery")
if err != nil {
  log.Error(err)
  return err
}

// inspect gallery struct
```

## View Your Subjects

This method returns a list of all subjects for a given gallery:

```go
subject, err := client.ViewSubject("MyGallery", "subjectID")
if err != nil {
  log.Error(err)
  return err
}

// inspect subject struct
```

## Remove a Subject

This method removes a subject from given gallery:

```go
subjectRequest := kairgo.RemoveSubjectRequest{
  GalleryName: "MyGallery",
  SubjectID: "subjectID",
}

result, err := client.RemoveSubject(&subjectRequest)

if err != nil {
  log.Error(err)
  return err
}

// Inspect result object for errors, statuses and messages
```

## Remove a Gallery

This method removes a given gallery:

```go
result, err := client.RemoveGallery("MyGallery")
if err != nil {
  log.Error(err)
  return err
}

// inspect result struct
```

## Enroll an Image

The **Enroll** method **registers a face for later recognitions**. Here's an example of enrolling a face (subject) using a method that accepts an image URL, and enrolls it as a new subject into your specified gallery:

```go
var (
  imageUrl    = "http://media.kairos.com/kairos-elizabeth.jpg"
  subjectID   = "Elizabeth"
  galleryName = "MyGallery"
)

result, err := client.Enroll(imageUrl, subjectID, galleryName, "", false)
if err != nil {
  log.Error(err)
  return err
}

// Inspect result object
```

## Recognize an Image

The **Recognize** method takes an image of a subject and **attempts to match it against a given gallery of previously-enrolled subjects**. Here's an example of recognizing a subject using a method that accepts an image URL, sends it to the API, and returns a match and confidence value:

```
var (
  imageUrl      = "http://media.kairos.com/kairos-elizabeth.jpg"
  galleryName   = "MyGallery"
  minHeadScale  = ".015"
  threshold     = "0.63"
  maxNumResults = 10
)

result, err := client.Recognize(imageUrl, galleryName, minHeadScale, threshold, maxNumResults)
if err != nil {
  log.Error(err)
  return err
}

// Inspect result object
```

## Detect Image Attributes

The **Detect** method takes an image of a subject and **returns various attributes pertaining to the face features**. Here's an example of using detect via method that accepts an image URL, sends it to the API, and returns face attributes:

```
var (
  detectRequest := kairgo.DetectRequest{
    Image: "http://media.kairos.com/kairos-elizabeth.jpg"
  }
)

result, err := client.Detect(&detectRequest)
if err != nil {
  log.Error(err)
  return err
}

// Inspect result object
```

## Verify image

The **Verify** method takes an image and verifies that it matches an existing subject in a gallery.  Here's an example of using verify via method that accepts a path to an image file, sends it to the API, and returns face attributes: 

```
result, err := client.ViewSubject("MyGallery", "subjectID")
if err != nil {
  log.Error(err)
  return err
}

// Inspect result object
```



---


## Copyrights

All rights belong [Kairos support page](http://www.kairos.com/support). 
