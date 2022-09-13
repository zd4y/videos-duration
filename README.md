# videos-duration

Get the total duration of one or more videos

## Installation

Clone this repository and enter:

```
go install .
```

##  Usage

The path of the videos is received from stdin, so for example you can get the total duration of all
the videos in the current directory like this:

```
ls | videos-duration
```

Or you could specify the paths with echo:

```
echo -e "video1.mp4\nvideo2.mp4" | videos-duration
```

Or simply running `videos-duration` and manually typing the paths and then using Ctrl+D

Or any other way of writing to stdin you like.

## Usage as a lf command

Add this to your `lfrc`:

```
cmd duration ${{
    total="$(printf '%s\n' $fx | videos-duration)"
    lf -remote "send $id echo '$total'"
}}
```

And now you can select the videos and use `:duration`
