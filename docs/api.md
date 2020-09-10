# Million Singer API Design

> Latest version of this document is on [hackmd](https://hackmd.io/aAbRR9mXSbmH0bOCpfE-RA)

## Columes of struct

### Game

- level IDs (array of integers)
- game ID (integer)

### Level

- Title (string)
- Song IDs (array of integers)

### Song
> TODO

**TODO: Add parameter & return**

```/api/auth```

- POST: Login
- DELETE: Logout

```/api/game/games```

- GET: Randomly get a game's info
- POST: Add a new game

```/api/game/games/${ID}```

- GET: Get a game's info
- DELETE: Delete the game (admin or author required)
- PATCH: Update the game (low priority)

```/api/game/games/search```

- GET: Create a search of games 

```/api/game/levels```

- POST: Add a new level

```/api/game/levels/${ID}```

- GET: Get a level's info
- DELETE: Delete the level (admin or author required)
- PATCH: Update the level (low priority)

```/api/game/songs/new```

**- POST: Add a new song**
```jsonld=
curl --location --request GET 'localhost/api/game/songs/1' \
--form 'file=@/home/james/Downlaods/[Toolbxs]Eminem - Beautiful (Edited) (Explicit)-English.srt' \
--form 'url=https://www.youtube.com/watch?v=avLxcVkPgug' \
--form 'language=en' \
--form 'name=Beautiful' \
--form 'singer=Eminem' \
--form 'genre=Hip-Hop,Rapping' \
--form 'miss_lyrics=30,31,32,33,34,35,36,37'
```

```/api/game/songs/${ID}```

- **GET: Get a song's info**
- DELETE: Delete the song (admin or author required)
- PATCH: Update the song (low priority)

~~```/api/game/songs/${ID}/lyrics```~~

~~- GET: List all lyrics info~~
~~- POST: Add new lyrics~~
~~- PATCH: Update lyrics (partly)~~