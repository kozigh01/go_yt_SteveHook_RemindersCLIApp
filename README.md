# Golang: youtube - Steve Hook - Reminders CLI App

youtube: [Reminders CLI App](https://www.youtube.com/playlist?list=PLsc-VaxfZl4cL9xE13tSe2Y9MOte9nSc8)
github: [my repo](https://github.com/kozigh01/go_yt_SteveHook_RemindersCLIApp)

## examples/cli-basis/os-args

* Sample command lines:

    ```
    $ go build
    $ ./os-args
    $ ./os-args aaa bbb ccc
    ```
## examples/cli-basics/flagset

* Sample command lines:

    ```
    $ go build
    $ ./flagset
    $ ./flagset abcd
    $ ./flagset help
    $ ./flagset greet
    $ ./flagset greet -msg "dude"
    $ ./flagset greet --msg "dude"
    $ ./flagset greet -msg="dude"
    $ ./flagset greet --msg="dude"
    $ ./flagset greet -help
    ```

## examples/cli-basics/flag-value

* Sample command lines:
    ```
    $ go build
    $ ./flag-value -id 11 --id 22 -id=33 --id=44
    ```