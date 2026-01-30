---
# How to build a Photo Database with Go and PocketBase
---

**Goal:** In this tutorial, you will learn how to create a simple photo database by combining the language Go and the Database Manager PocketBase.

I will assume you have basic coding knowledge and know the basics of Go. If you don't have any knowledge of PocketBase, that's okay! PocketBase is one of the easiest database managers I have worked with so far. By the end of this tutorial, you will have a working app where you can upload and view the images you uploaded.

## Checklist
- Basic knowledge of Go and HTML
- Go installed on your machine
- A code editor (I use Visual Studio Code)

## 1. Project Setup
The first thing to do is to create a folder named `photodb`. Open this folder in your terminal and run the following command to initialize your Go project:

```
go mod init photodb
```

Next, you need to download the PocketBase libary. This allows your Go code to talk to the database engine. Run this command:

```
go get github.com/pocketbase/pocketbase
```

For the last step, create a subfolder named `pb_public` inside your `photodb` folder. THis is where your HTML and CSS files will live.
Now create a File in the `photodb` folder named `main.go`, this will be where the code will live.


After those first few steps, your structure should look like this:
```
photodb/
|--- main.go             --> Main Code
|--- go.mod
|--- pb_public/          --> Folder where your Frontend will be
      |--- index.html
      |--- styles.css
```

## 2. The Backend with Go
PocketBase is usually a standalone program, but since it is written in Go, we can import it directly into our code, which makes a lot of things easier. It also allows us to serve our HTML files from the same server.
Now write the following code in the `main.go` file.

```
package main

import (
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {

		e.Router.GET("/{path...}", func(e *core.RequestEvent) error {
			fs := http.FileServer(http.Dir("./pb_public"))
			
			fs.ServeHTTP(e.Response, e.Request)
			
			return nil
		})

		records, err := app.FindRecordsByFilter("photos", "1=1", "-created", 100, 0)
		if err != nil {
			log.Println("Fehler beim Abrufen der Bilder:", err)
		} else {
			log.Println("----- STARTUP BILD-CHECK -----")
			if len(records) == 0 {
				log.Println("Noch keine Bilder in der Datenbank.")
			}
			for _, r := range records {
				log.Printf("Foto gefunden: %s (ID: %s)", r.GetString("title"), r.Id)
			}
			log.Println("------------------------------")
		}

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
```

## The Frontend with HMTL / CSS
To make the Photo Database more accessible, we will now create a HTML File. This File should live in the `pb_public` folder. 
We will use standart HTML form. But Notice the `enctype="multipart/form-data"` attribute in the form tag. Without this, the browser will only send the filename, not the actual image file.!
```
<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="UTF-8">
    <title>Foto Upload</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <h1>Neues Foto hochladen</h1>

    <form action="/api/collections/photos/records" method="post" enctype="multipart/form-data">       --> Upload images
        
        <div>
            <label for="title">Titel des Bildes:</label>
            <input type="text" id="title" name="title" required>
        </div>

        <div>
            <label for="image">Datei auswählen:</label>
            <input type="file" id="image" name="image" required>
        </div>

        <button type="submit">Hochladen</button>
    </form>

    <form action="/api/collections/photos/records" method="get" style="margin-top:20px;">            --> View All images
        <button type="submit">Alle Fotos anzeigen</button>
    </form>

</body>
</html>
```
If you want, you can add a small CSS File, all I did was this:
```
        body { 
            font-family: sans-serif; padding: 2rem; 
        }
        
        form { 
            border: 1px solid #ccc; padding: 1rem; width: 300px; 
        }

        div { 
            margin-bottom: 1rem; 
        }

        label { 
            display: block; margin-bottom: .5rem; 
        }
```

## 4. Setting up the Database
This is a very important step, because without this step, you won't have a Database where you can store your images. 
1. Start the server with running this command in your terminal:
   ```
   go run main.go serve
   ```
3. In your terminal should come up three links looking like this:

   <img width="410" height="64" alt="image" src="https://github.com/user-attachments/assets/ab4df78f-ca70-4fe5-9edb-efbfe5ec13ad" />
   
   Open the "Dashboard" Link in your Broswer.
5. Create an Account with email and password.
6. Create a collection:
  - Click on "New Collection"
  - Name it "photos"
  - Click on "New field" and select **text**. Name this field `title`.
  - Create another field and select the **file** type. Name it `photo`.
5. Now make sure to have your API Rules empty, or else it won't allow you to see your files when using your Photo Database. If you did it correct, it should look like this:
<img width="666" height="523" alt="image" src="https://github.com/user-attachments/assets/fba0c218-4031-43bd-a6e1-622746e3d23d" />

   Click on save and your done with setting up your Database with PocketBase.

## Testing
Now, if you did everything right, you should be able to click on the first of the three links you got while running the `run` command. If you open this link in your Browser you should see your HTML File.
Have Fun trying out!

### GIF
Here is a GIF, which can help you to visualize everything how it should look.
![Finished_Project](https://github.com/user-attachments/assets/f4076dfe-53ef-441b-a90f-d9ba446a41f5)

## Possible Troublemakers
It is really important to observe the API-Rules of your collection. Because I had the problem, that I kept the standart settings. As a result I kept getting the message: `"Only superusers can perform this action.", "status": 403` when trying to upload a picture.
So if you get this Error, go to your Collection and make sure the API-Rules are empty.

