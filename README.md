# whatsapp-analyzer

It helps to analyze your Whatsapp chat records with making parse on chat texts. As example you can measure the aggression, you can see the words in order of use and you can see the frequency of messaging.


# Getting Started

1. First download the app which is available for your operating system from here; https://drive.google.com/open?id=0B2ClpaG8PdV6WHNjZ0pLVmhaTVU

2. After, add the folder of app to your PATH.

    In OSX or any of Linux you can do it in terminal with this command:
    ```sh
    export PATH = "$PATH:$HOME/folderofapp"
    ```

    For Windows follow these instructions:
    1. From the desktop, right click the Computer icon.
    2. Choose Properties from the context menu.
    3. Click the Advanced system settings link.
    4. Click Environment Variables. In the section System Variables, find the PATH environment variable and select it.
    5. Click Edit. If the PATH environment variable does not exist, click New.
    6. In the Edit System Variable (or New System Variable) window, specify the value of the PATH environment variable. Click OK. Close all remaining windows by clicking OK.

<br>

3. You must send the chat record without media to your email from Whatsapp. After export, download the chat record from your email.
<br>For this: https://www.whatsapp.com/faq/en/android/23756533

# Usage

Open your command line in directory which you was installed the program. Now you are ready to use these commands.

| Commands  | Description                                                                |
|-----------|----------------------------------------------------------------------------|
| file      | For read the chat record you must send the file address with this command. While alone it returns a list of most used words.                                                                 |
| username  | You can make querying by username in Whatsapp chat.                        |
| limit     | You can limit the list from start to sended value.                         |
| start     | It sets starting index of list to sended value.                            |
| word      | Finds the count of specific word or sentence.                              |
| negatives | You can observe the aggression with identifying some aggressive words.     |
| messagef  | It provides a graph for observe messaging frequency.                       |
| timef     | It provides a graph for observe messaging frequency in time periods.       |
| printf    | While with <div style="color:orange;display:inline;">messagef</div> it sorts messaging frequency by activity and prints. While with <div style="color:orange;display:inline;">timef</div> it prints messaging frequency in time periods.|


# Examples
For find the most used words in general;
```sh
analyzer --file "C:\chatrecord.txt"
```

For find the most used words for specific person;
```sh
analyzer --file "C:\chatrecord.txt" --username "Cem Asma"
```

If you do not send limit, program sets it defaultly 10. For see more or less words;
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
analyzer --file "C:\chatrecord.txt" --start 10 --limit 20
```

For find the count of specific word or sentence;
```sh
analyzer --file "C:\chatrecord.txt" --word "test"
```

```sh
analyzer --file "C:\chatrecord.txt" --word "hello world"
```

If you send the negative words in a file via file address you can measure the aggression by these words.

This file's content must be like this;
```
fuck you
fuck
shit
```

After creation this file you can use it like this;
```sh
analyzer --file "C:\chatrecord.txt" --negatives "C:\negatives.txt"
```

It counts your messages in day  by day and draws a graphs. It saves the graph image on where you are. With that you can observe messaging frequency;
```sh
analyzer --file "C:\chatrecord.txt" --messagef
```

```sh
analyzer --file "C:\chatrecord.txt" --username "Cem Asma" --messagef
```

It sorts message frequency by message counts and prints.
```sh
analyzer --file "C:\chatrecord.txt" --messagef --printf
```

It provides a graph for observe messaging frequency in time periods.
```sh
analyzer --file "C:\chatrecord.txt" --timef
```

It prints the values of messaging frequency by time periods.
```sh
analyzer --file "C:\chatrecord.txt" --timef --printf
```