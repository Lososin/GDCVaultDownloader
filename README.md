#GDC Vault m3u8 file downloader

####Help
go run ./main.go -d

  -path string
        Path to save video
  -queue string
        file, that contains names and urls
  -url string
        url to m3u8 file

####Download single file
go run ./main.go -url="https://url/to/master.m3u8" -path="path/to/save/file.mkv"

####Download multiple files
go run ./main.go -queue="path/to/url_list.txt"

####Url_List.txt format
PathToSave1
url1
pathToSave2
url2
...

List.txt as template included in project.

####How to find url to m3u8?
1. Open code of page.
2. Run video.
3. Open Network tab, and find master.m3u8 in names list.

![Alt text](/Screenshots/1.jpg?raw=true "m3u8 location")