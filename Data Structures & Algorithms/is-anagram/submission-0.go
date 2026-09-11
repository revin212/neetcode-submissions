func isAnagram(s string, t string) bool {
	if(len(s) != len(t)){
		return false;
	}

	var charCount1 = make(map[byte]int)
	var charCount2 = make(map[byte]int)

	for i:=0; i < len(s); i++ {
		charCount1[s[i]] ++;
		charCount2[t[i]] ++;
	}

	for key, value := range charCount1 {
    	if(charCount2[key] != value){
			return false;
		}
	}

	return true;
}
