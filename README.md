# awsm

Go package for AWS tasks

## Setup
* Create an IAM user with Access Keys
* Install AWS CLI (https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* Configure AWS credentials:
    - run `aws configure`
    - Set AccessKey and SecretAccessKey

## Upload Config 
```
type UploadConfig struct{
  Profile     string 
  Region      string 
  Bucket      string 
  FilePath    string 
  BucketPath  string 
  ACL         ObjectCannedACL 
  ContentType string
}
```

## Upload URL 
```
cfg := &UploadConfig{...} 
cfg.PublicURL() 
// https://(bucket).s3.(region).amazonaws.com/(bucketPath)
```

## Upload File 
```
err := awsm.UploadFile(*UploadConfig)
```