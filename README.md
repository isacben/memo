# memo

Organize key-value pair information and links in categories in a local HTML you can load in your browser.

## Usage

1. Make a copy the file `content.json.example` and name it `content.json`.
2. Edit the `content.json` file; add sections as needed.
3. Run `go run main.go` to render the html with the content. A file called `memo.html` will be created.
4. Open the `memo.html` file in your web browser.

## Update the content

1. Update the `content.json` file as needed.
2. Then run `go run main.go` to update the `memo.html` output file.
3. Refresh the `memo.html` page in the browser to see the updated content.

## Example of the `content.json` file

You can add as many sections as needed.

```json
{
  "sections": [
    {
      "name": "Section 1",
      "items": [
        {
          "type": "text",
          "description": "Some ID",
          "value": "12345-12345"
        },
        {
          "type": "link",
          "description": "Some URL",
          "value": "https://google.com"
        }
      ]
    }
  ]
}
```

--  
@isacben
