# whatsapp-analyzer

It find out most used words in your Whatsapp chat records.

# Installation

* For Linux
[downloadadress](downloadlinuxlink)
* For OSX
[downloadadress](downloadlinuxlink)
* For Windows
[downloadadress](downloadlinuxlink)

# Usage

1. First you must send the chat record to your email from Whatsapp.
    For this: https://www.whatsapp.com/faq/en/s60/21055276
    
    Download the chat record from your email after export.

<br>
2. Open your command line in directory where is you was installed the program. Now you can use these commands.
* file
* username
* limit
* start

For find most used words in general;
```sh
analyzer --file "C:\chatrecord.txt"
```

For find most used words for specific person;
```sh
analyzer --file "C:\chatrecord.txt" --username "Cem Asma"
```

If you does not send limit, program sets it defaultly 10. For see more or less words;
```sh
analyzer --file "C:\chatrecord.txt" --limit 20
```

```sh
analyzer --file "C:\chatrecord.txt" --username "Cem Asma" --limit 20
```

For set the starting index;
```sh
analyzer --file "C:\chatrecord.txt" --start 10
```

You can combine start and limit;
```sh
analyzer --file "C:\chatrecor.txt" --start 10 --limit 20
```