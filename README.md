# whatsapp-analyzer

It helps to analyze your Whatsapp chat records with making parse on records. As example you can measure the aggression, you can see the words in order of use and you can see the frequency of messaging.


# Getting Started

1. First build the project by yourself or download the app which is available for your operating system from releases:

    https://github.com/cemasma/whatsappanalyzer/releases/latest

2. After, add the folder of app to your PATH.

    In OSX or any of Linux you can do it in terminal with this command:
    ```sh
    export PATH = "$PATH:$HOME/folderofapp"
    ```

    Follow these instructions for Windows:
    1. From the desktop, right click the Computer icon.
    2. Choose Properties from the context menu.
    3. Click the Advanced system settings link.
    4. Click Environment Variables. In the section System Variables, find the PATH environment variable and select it.
    5. Click Edit. If the PATH environment variable does not exist, click New.
    6. In the Edit System Variable (or New System Variable) window, specify the value of the PATH environment variable. Click OK. Close all remaining windows by clicking OK.

<br>

3. You must send the chat record without media to your email from Whatsapp. After export, download the chat record from your email.
<br>For this: https://www.whatsapp.com/faq/en/android/23756533

# Build Project

1. Build the project via Go.
```sh
go build main.go
```

2. I renamed the executable as "analyzer" and I add it to paths. So I can run it anywhere. You can skip this option but remember that if you want to run project then you must be in the directory which is contains the executable.

```sh
analyzer --help
```

# Usage

After those steps you are ready to use the app in console. I suggest the help command for the first opening.

```sh
analyzer --help
```

| Commands  | Description                                                                |
|-----------|----------------------------------------------------------------------------|
| help      | It shows all the commands with their abilities.                            |
| file      | For read the chat record you must send the file address with this command. While alone it returns a list of most used words.                                                                 |
| username  | You can make querying by username in Whatsapp chat.                        |
| limit     | You can limit the list from start to sended value.                         |
| start     | It sets starting index of list to sended value.                            |
| word      | Finds the count of specific word or sentence.                              |
| aggression| You can observe the aggression by users.                                   |
| topic     | It provides topics that spoken in chat and filters that by date.           |
| messagef  | It provides a graph for observe messaging frequency.                       |
| timef     | It provides a graph for observe messaging frequency in time periods.       |
| printf    | While with messagef it sorts messaging frequency by activity and prints. While with timef it prints messaging frequency in time periods.|


## Examples

* For find the most used words in general.
    ```sh
    analyzer --file "chatrecord.txt"
    ```

* For find the most used words for specific person.
    ```sh
    analyzer --file "chatrecord.txt" --username "Cem Asma"
    ```

* If you do not send limit, program sets it defaultly 10. For see more or less words.
    ```sh
    analyzer --file "chatrecord.txt" --limit 20
    ```

    ```sh
    analyzer --file "chatrecord.txt" --username "Cem Asma" --limit 20
    ```

* For set the starting index.
    ```sh
    analyzer --file "chatrecord.txt" --start 10
    ```

* You can combine start and limit.
    ```sh
    analyzer --file "chatrecord.txt" --start 10 --limit 20
    ```

* For find the count of specific word or sentence.
    ```sh
    analyzer --file "chatrecord.txt" --word "test"
    ```

    ```sh
    analyzer --file "chatrecord.txt" --word "hello world"
    ```

* If you send the username and specify the aggression parameters as true you can measure the users aggression with artificial neural network.

    ```sh
    analyzer --file "chatrecord.txt" --username "Cem Asma" --agression true
    ```

* Also we can observe the topics that spoken in chat by the date information.

    ```sh
    analyzer --date "5.30.20" --topic true
    ```

* It counts your messages in day  by day and draws a graphs. It saves the graph image on where you are. With that you can observe messaging frequency.
    ```sh
    analyzer --file "chatrecord.txt" --messagef
    ```

    ```sh
    analyzer --file "chatrecord.txt" --username "Cem Asma" --messagef
    ```

* It sorts message frequency by message counts and prints.
    ```sh
    analyzer --file "chatrecord.txt" --messagef --printf
    ```

* It provides a graph for observe messaging frequency in time periods.
    ```sh
    analyzer --file "chatrecord.txt" --timef
    ```

* It prints the values of messaging frequency by time periods.
    ```sh
    analyzer --file "chatrecord.txt" --timef --printf
    ```
