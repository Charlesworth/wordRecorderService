# wordRecorderService
This application takes in sentences, stores the component words and can output statistics about them, such as the most frequent letters.

### Running

##### Running binary
I've supplied a windows executable binary. Please download and open the terminal in the same directory as the binary and run:

    $ ./wordRecorderService.exe

##### Running from source
First have the [go programming runtime installed](https://golang.org/). Next download the source and cd into its root directory.

    $ go build
Then on Linux:

    $ ./webCrawler
Or Windows:

    $ ./webCrawler.exe

#### User Flags
- "-inputPort=[port number here]" use input port to configure which tcp port should be used for sentence input. Default it 5555.
- "-statsPort=[port number here]" use stats port to configure which tcp port should be used for statistic output. Default it 8080.
- "-h" shows flag information

### API

#### - Put a sentence:
##### input port (default ':5555') `PUT` /
Add a sentence of words to the wordRecorderService. Please include the sentence in the request body.

#### - Get Statistics:
##### stats port (default ':8080') `GET` /stats
Returns statistics in JSON format on:
- word count
- most frequent 5 words
- most frequent 5 letters

### Testing

To run tests, in the root directory of the project please run:

    $ go test ./... -v -cover
