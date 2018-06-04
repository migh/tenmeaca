# Tenmeacá

A go mailing platform in very early development process.
*****
## Motivation

There are already several services for sending both marketing and transactional emails, also there are already a few different libraries for creating and sending mails in any language you can think of. Despite that fact, it is very hard to get started, I guess the main reason is most people is not even aware of the importance of emails and all the process involved in creating a nice and clean, successful email.

This library aims to simplify the email creation pipeline, a little more detail will be given in the roadmap, the main aspect being to focus on the user's point of view flow instead of the developer's point of view.  

## Roadmap

1. Use TOML for interacting with the application.
2. Follow a component approach to template creation.
3. Use Webpack as module bundler for template creation.
4. Use nested templates for separating concerns, like different stages in a process.
5. Create a stand-alone service that can be used inside a Docker container.

## installation

It is still too early to talk about that.

We use [Hermes](https://github.com/matcornic/hermes) as Mail Generator.
Here are a few suggestion to look for templates:
* [Postmark Transactional Email Templates](https://github.com/wildbit/postmark-templates)
* [EmailOctopus Templates](https://github.com/threeheartsdigital/emailoctopus-templates)
