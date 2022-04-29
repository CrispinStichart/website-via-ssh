# My Blog, via SSH

There's a cool project called [Charm](https://charm.sh/) that provides various libraries for making terminal user interfaces (TUIs) easy to write and pleasant to use. Charm's BubbleTea library the terminal equivalent of a GUI toolkit like GKT or Qt.

Charm also provides Wish, an easy way to make your terminal apps available over SSH.

The idea was so neat I had to try it myself, so I started learning Go and wrote a program that, using BubbleTea, the Glamour markdown renderer, and Wish, will serve up my blog via SSH.

It's my first ever Go program, so don't think too badly of me if there's un-idiomatic code.

# So, like, how do I SSH your blog, bro?

Although this program works as advertised, I don't actually have it running on a server anywhere. Eventually I'm going to start renting a server to do fun things with, and I'll stick it on there, but right now I'm trying to save money until I land a job.

# Thoughts on Charm

It's pretty cool! The BubbleTea documentation could stand to go a little bit further in depth about how it works, but at least they provide a whole bunch of example programs. If I ever need to make an advanced TUI in the future, I will definitely consider it.

The Wish component that provides SSH was dead simple, I just copy and pasted a bit of code and it Just Worked.