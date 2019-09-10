#!/bin/bash

for paper in {0..196}
do
	echo "Downloading paper $paper"

	# Write Text 
	curl -s --header "Content-Type: application/json" --request POST \
		--data '{"language":"en", "paper":'$paper', "plaintext":true}' http://urantiabook.global.mac/paper \
		| jq '.paper.sections' \
		| jq '.[].text' > paper$paper.txt

	# echo "  generating paper$paper.aiff"
	# say -v Alex -f paper$paper.txt -o paper$paper.aiff
	# echo "  converting paper$paper.aiff --> paper$paper.mp3"
	# lame -m m paper$paper.aiff paper$paper.mp3

	# Remove unnecessary files
	# rm paper$paper.aiff


	let "section=0"
	 
 	# Read one section per line from file
	while IFS= read -r line
	do
		echo "$line" > "paper$paper-$section.txt"
		
		echo "    generating paper$paper-$section.aiff"
		say -v Alex -f paper$paper-$section.txt -o paper$paper-$section.aiff
		echo "    converting paper$paper-$section.aiff --> paper$paper-$section.mp3"
		lame -m m paper$paper-$section.aiff paper$paper-$section.mp3

		# Remove unnecessary files
		rm paper$paper-$section.txt
		rm paper$paper-$section.aiff

		let "section=section+1"
	done < paper$paper.txt

	rm paper$paper.txt
	 
done
