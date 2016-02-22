#!/bin/bash
infoFile=$SR_CODE_BASE/reltools/codegentools/._genInfo/generatedGoFiles.txt
#sort $infoFile|
touch $infoFile
uniq $infoFile >$infoFile.tmp
for srcFile in `cat $infoFile.tmp`;
do
	 touch $srcFile
	 rm $srcFile
done    
rm $infoFile
rm $infoFile.tmp

