*First attempt*
Approach:
Loop through first string and create a char count map and then loop through second string to decrease each count if not found or <0 means not a same word 
Thought on Big-O: O(n)
Result:
Pass with 20.33 percentile (Not a result I hope for)

*Second attempt*
Approach:
Searched on others submission and found that with the same idea
runeSet[char]++ is much faster than runeSet[char] = runeSet[char]+1
Result:
Pass with 92.15 percentile
