# whatsapp-analyzer

It helps to analyze your Whatsapp chat records with making parse on chat texts. As example you can measure the aggression, you can see the words in order of use and you can see the frequency of talking.


# Installation

* For Linux
[downloadadress](downloadlinuxlink)
* For OSX
[downloadadress](downloadlinuxlink)
* For Windows
[downloadadress](downloadlinuxlink)

# Usage

First you must send the chat record to your email from Whatsapp. After export, download the chat record from your email.
<br>For this: https://www.whatsapp.com/faq/en/s60/21055276

Open your command line in directory where you was installed the program. Now you can use these commands.

* file
* username
* limit
* start
* word
* negatives

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

For find the count of specific word or sentence;
```sh
analyzer --file "C:\chatrecor.txt" --word "test"
```

```sh
analyzer --file "C:\chatrecor.txt" --word "hello world"
```

If you send the negative words in a file you can measure the aggression by these words.

That file's content must be like this;
```
negativeword1
negativeword2
negativeword3
```

After creation that file you can use it like this;
```sh
analyzer --file "C:\chatrecord.txt" --negatives "C:\negatives.txt"
```

For a specific user;
```sh
analyzer --file "C:\chatrecord.txt" --username "Cem Asma" --negatives "C:\negatives.txt"
```

