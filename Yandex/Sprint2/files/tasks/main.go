k

func ModifyFile(filename string, pos int, val string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(int64(pos), 0)
	file.WriteString(val)
}

func ReadContent(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	content, err := io.ReadAll(bufio.NewReader(file))
	if err != nil {
		return ""
	}

	return string(content)
}

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	inputFile.Seek(int64(startpos), 0)
	content, err := io.ReadAll(bufio.NewReader(inputFile))
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	_, err = outputFile.Write(content)
	if err != nil {
		return err
	}

	return nil
}
