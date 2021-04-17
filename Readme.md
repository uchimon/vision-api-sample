# Outline

 - Sample code for Google Vision API
 - This calls Google Vision API and get info of `LANDMARK` and `LABEL`
 - Read list of targe image uri from `./imglist.txt`
 - Write output json to `./out.json`

# Setup

 1. set up GCP project. please check Google Cloud document.
 2. set `GOOGLE_APPLICATION_CREDENTIALS`
 ```
 export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/my-key.json"
 ```
 3. put list file of target image uri
 ```
 https://xxxxx/image.jpg
 https://xxxxx/image2.jpg
 ```
 4. run main.go
 ```
 go run main.go
 ```

 # Out put json file sample

 ```
[
   {
      "image":"https://xxxxx/image.jpg",
      "labels":[
         {
            "description":"Sky",
            "score":0.9676917,
            "topicality":0.9676917
         },
         {
            "description":"Billboard",
            "score":0.9046523,
            "topicality":0.9046523
         }
      ],
      "landmarks":[
         "Dotombori Glico Sign"
      ]
   }
]
```
