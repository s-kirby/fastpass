# FastPass

(it's not like the disney thing)

FastPass or `fp` intends to be the quickest password manager.

It uses fuzzy searching and learns your password access tendencies to retrieve a password with as few keystrokes as possible.

By default it generates easy to remember passwords using human words.

## Example

I don't want to have to type this 15 times a day:

```bash
pass -c pornhub.com
```

with fp I just type

```bash
$ fp p
other matches: ammarb36@pornhub.com ammarb36@pornhut.com ammarb36@papajohns.com ammarb36@pizzahut.com ammarb36@paypal.com 
ammarb36@pornhub.com -> "ViolentRelativeDamsPreferences441" Copied!
```

## Install

```bash
go get -u github.com/s-kirby/fastpass/cmd/fp
```

## Features 

- Encryption
- Fuzzy searching
- Notes
- Ranking based on access frequency
- Password and key file support
- Multiple password generation strategies

## Generators

### Human

The human password generator uses the lists in world_list/ to generate passwords.

It uses the following format: `<Adjective><Adjective><Noun><Noun><random num 000->999>`

It uses about 55 bits of entropy.

### Hex

Hex generates 16 random hex digits.

It uses 64 bits of entropy.

### Base62

Base62 generates 16 random base62 digits.

It uses 96 bits of entropy.

## Password caching

fp caches secrets after an open in `/dev/shm/fp.secret`