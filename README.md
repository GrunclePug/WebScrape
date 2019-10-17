# WebScrape

A web scraper built in Go

## What is it?

<p>A web scraper allows the user to 'scrape' data from any webpages on the internet,</p>
<p>In this case, either whole webpages, or the html tag of their choosing.</p>

## Usage
The user can either provide the html tag they want, or they can omit it.
Omitting any html tags will print the entire webpage.

<b>Example:</b>
the input `https://example.com` would print the entire webpage's content

Here is an example of using WebScrape with a provided tag:
```go
Go Web Scraper v1.0
Created by: Chad Humphries (GrunclePug)
---------------------------------------------------------------------------------------------------------
Input a website URL
$> div https://example.com
Result:
<div>
    <h1>Example Domain</h1>
    <p>This domain is established to be used for illustrative examples in documents. You may use this
    domain in examples without prior coordination or asking for permission.</p>
    <p><a href="http://www.iana.org/domains/example">More information...</a></p>
</div>
New location:
$> quit
exit status 1
```

## Requirements

Go 1.13.1 or above (may work on lower versions, but officially unsupported).

## Author

Chad Humphries - <a href="https://grunclepug.com/" title="Go to my website">My Website</a>
<p>Check out some of my other projects: <a href="https://github.com/GrunclePug?tab=repositories" title="Other projects of mine">here</a>
