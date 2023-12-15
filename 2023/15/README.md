## \-\-\- Day 15: Lens Library ---

The newly-focused parabolic reflector dish is sending all of the collected light to a point on the side of yet another mountain - the largest mountain on Lava Island. As you approach the mountain, you find that the light is being collected by the wall of a large facility embedded in the mountainside.

You find a door under a large sign that says "Lava Production Facility" and next to a smaller sign that says "Danger - Personal Protective Equipment required beyond this point".

As you step inside, you are immediately greeted by a somewhat panicked reindeer wearing goggles and a loose-fitting [hard hat](https://en.wikipedia.org/wiki/Hard_hat). The reindeer leads you to a shelf of goggles and hard hats (you quickly find some that fit) and then further into the facility. At one point, you pass a button with a faint snout mark and the label "PUSH FOR HELP". No wonder you were loaded into that [trebuchet](1) so quickly!

You pass through a final set of doors surrounded with even more warning signs and into what must be the room that collects all of the light from outside. As you admire the large assortment of lenses available to further focus the light, the reindeer brings you a book titled "Initialization Manual".

"Hello!", the book cheerfully begins, apparently unaware of the concerned reindeer reading over your shoulder. "This procedure will let you bring the Lava Production Facility online - all without burning or melting anything unintended!"

"Before you begin, please be prepared to use the Holiday ASCII String Helper algorithm (appendix 1A)." You turn to appendix 1A. The reindeer leans closer with interest.

The HASH algorithm is a way to turn any [string](https://en.wikipedia.org/wiki/String_(computer_science)) of characters into a single _number_ in the range 0 to 255. To run the HASH algorithm on a string, start with a _current value_ of `0`. Then, for each character in the string starting from the beginning:

- Determine the [ASCII code](https://en.wikipedia.org/wiki/ASCII#Printable_characters) for the current character of the string.
- Increase the _current value_ by the ASCII code you just determined.
- Set the _current value_ to itself multiplied by `17`.
- Set the _current value_ to the [remainder](https://en.wikipedia.org/wiki/Modulo) of dividing itself by `256`.

After following these steps for each character in the string in order, the _current value_ is the output of the HASH algorithm.

So, to find the result of running the HASH algorithm on the string `HASH`:

- The _current value_ starts at `0`.
- The first character is `H`; its ASCII code is `72`.
- The _current value_ increases to `72`.
- The _current value_ is multiplied by `17` to become `1224`.
- The _current value_ becomes `200` (the remainder of `1224` divided by `256`).
- The next character is `A`; its ASCII code is `65`.
- The _current value_ increases to `265`.
- The _current value_ is multiplied by `17` to become `4505`.
- The _current value_ becomes `153` (the remainder of `4505` divided by `256`).
- The next character is `S`; its ASCII code is `83`.
- The _current value_ increases to `236`.
- The _current value_ is multiplied by `17` to become `4012`.
- The _current value_ becomes `172` (the remainder of `4012` divided by `256`).
- The next character is `H`; its ASCII code is `72`.
- The _current value_ increases to `244`.
- The _current value_ is multiplied by `17` to become `4148`.
- The _current value_ becomes `52` (the remainder of `4148` divided by `256`).

So, the result of running the HASH algorithm on the string `HASH` is `52`.

The _initialization sequence_ (your puzzle input) is a comma-separated list of steps to start the Lava Production Facility. _Ignore newline characters_ when parsing the initialization sequence. To verify that your HASH algorithm is working, the book offers the sum of the result of running the HASH algorithm on each step in the initialization sequence.

For example:

```
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
```

This initialization sequence specifies 11 individual steps; the result of running the HASH algorithm on each of the steps is as follows:

- `rn=1` becomes `30`.
- `cm-` becomes `253`.
- `qp=3` becomes `97`.
- `cm=2` becomes `47`.
- `qp-` becomes `14`.
- `pc=4` becomes `180`.
- `ot=9` becomes `9`.
- `ab=5` becomes `197`.
- `pc-` becomes `48`.
- `pc=6` becomes `214`.
- `ot=7` becomes `231`.

In this example, the sum of these results is `1320`. Unfortunately, the reindeer has stolen the page containing the expected verification number and is currently running around the facility with it excitedly.

Run the HASH algorithm on each step in the initialization sequence. _What is the sum of the results?_ (The initialization sequence is one long line; be careful when copy-pasting it.)