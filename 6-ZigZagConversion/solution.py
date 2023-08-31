'''
    [6] Zig Zag Conversion
    [Difficulty] Medium
    The string "PAYPALISHIRING" is written in a zigzag pattern on a given
    number of rows like this:

        P   A   H   N
        A P L S I I G
        Y   I   R

    And then read line by line: "PAHNAPLSIIGYIR"

    Write the code that will take a string and make this conversion
    given a number of rows:

    string convert(string s, int numRows);

    [Runtime] 55ms - Better than 89.37%
    [Memory]  16.28MB - Better than 96.81%
'''


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        word = [''] * numRows

        down = True
        line = 0
        if (numRows == 1):
            return s
        for i in s:
            word[line] = word[line] + i
            if (down):
                line += 1
                if (line >= numRows):
                    line -= 2
                    down = False
            else:
                line -= 1
                if (line < 0):
                    down = True
                    line += 2
        toPrint = ""
        for i in word:
            toPrint += str(i)
        return toPrint
