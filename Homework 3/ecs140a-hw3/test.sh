#!/bin/bash

SOLUTION=hw3.l  # only need to configure this if necessary
TESTFILE=test.l
TMPFILE=check.l
REFOUT=Output.correct
TMPOUT=Output.your

if ! type "clisp" &> /dev/null; then
    echo "clisp is not available."
    echo "If you are on Kodethon, make sure to change your execution environment."
    echo "If you are on your own machine, make sure to install CLisp."
    exit 1
fi

cat > $TMPFILE <<EOF
(load "$SOLUTION")
(load "$TESTFILE")
EOF
	if [[ $1 == '1' ]]; then
		echo "(tec)" >> $TMPFILE
		echo "(tmmm)" >> $TMPFILE
		echo "(tc)" >> $TMPFILE
		echo "(tp)" >> $TMPFILE
		echo "(tbl)" >> $TMPFILE
		echo "(tq)" >> $TMPFILE
		REFOUT=OutputPartOne.correct
	else
		if [[ $1 == '2' ]] ; then
			echo "(tm)" >> $TMPFILE
			REFOUT=OutputPartTwo.correct
		else
			if [[ $1 == '3' ]]; then
				echo "(tma)" >> $TMPFILE
				echo "(tmm)" >> $TMPFILE
				echo "(tmt)" >> $TMPFILE
				REFOUT=OutputPartThree.correct
			else
				echo "(test)" >> $TMPFILE
				REFOUT=Output.correct
			fi

		fi

	fi

# cleanup
echo "(1) remove existing output file"
rm -f $TMPOUT

# run
echo "(2) load your solution and test"
clisp $TMPFILE > $TMPOUT

# diff
echo "(3) check for differences"
diff -s $TMPOUT $REFOUT > /dev/null
if [ $? -eq 0 ]; then
    echo "The files are identical"
    #rm -f $TMPOUT
    rm -f $TMPFILE
 else
 	rm -f $TMPFILE
 	diff $TMPOUT $REFOUT
 fi
