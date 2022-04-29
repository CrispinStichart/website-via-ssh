# My Blog, via SSH

There's a cool project called [Charm](https://charm.sh/) that provides various libraries for making terminal user interfaces (TUIs) easy to write and pleasant to use. Charm's BubbleTea library the terminal equivalent of a GUI toolkit like GKT or Qt.

Charm also provides Wish, an easy way to make your terminal apps available over SSH.

The idea was so neat I had to try it myself, so I started learning Go and wrote a program that, using BubbleTea, the Glamour markdown renderer, and Wish, will serve up my blog via SSH.

It's my first ever Go program, so don't think too badly of me if there's un-idiomatic code.

## Current Status

No SSH yet.

It displays all my blog posts, allows you to select and read individual posts, and then back out to the list of posts again.

to do:
 - [X] integrate Wish for SSH support
 - [ ] parse the titles for the posts in the posts list
 - [ ] sort the posts