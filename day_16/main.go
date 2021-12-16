package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Packet struct {
	version    int64
	typeID     int64
	content    string
	lengthType int
	length     int64
	subPackets []Packet
}

func main() {
	transmission := parseInputs(getInputsByLine())

	pkt, _ := readPacket(transmission, 0, true)

	fmt.Printf("Solution 1: %d\n", sumVersionNumbers(pkt))
	fmt.Printf("Solution 2: %d\n", operate(pkt))
}

func operate(pkt Packet) int64 {
	if pkt.typeID == 4 {
		br := 0
		number := ""
		for br < len(pkt.content) {
			number = fmt.Sprintf("%s%s", number, pkt.content[br+1:br+5])
			br += 5
		}
		res, _ := strconv.ParseInt(number, 2, 64)

		return res
	}

	res := int64(0)
	switch pkt.typeID {
	case 0: //sum
		for _, spkt := range pkt.subPackets {
			res += operate(spkt)
		}
	case 1: //product
		res = operate(pkt.subPackets[0])
		for i := 1; i < len(pkt.subPackets); i++ {
			res = res * operate(pkt.subPackets[i])
		}
	case 2: //min
		res = operate(pkt.subPackets[0])
		for i := 1; i < len(pkt.subPackets); i++ {
			nv := operate(pkt.subPackets[i])
			if res > nv {
				res = nv
			}
		}
	case 3: //max
		res = operate(pkt.subPackets[0])
		for i := 1; i < len(pkt.subPackets); i++ {
			nv := operate(pkt.subPackets[i])
			if res < nv {
				res = nv
			}
		}
	case 5: //gt
		res = 0
		if operate(pkt.subPackets[0]) > operate(pkt.subPackets[1]) {
			res = 1
		}
	case 6: //lt
		res = 0
		if operate(pkt.subPackets[0]) < operate(pkt.subPackets[1]) {
			res = 1
		}
	case 7: //eq
		res = 0
		if operate(pkt.subPackets[0]) == operate(pkt.subPackets[1]) {
			res = 1
		}
	}

	return res
}

func sumVersionNumbers(pkt Packet) int64 {
	if pkt.typeID == 4 {
		return pkt.version
	}

	sum := pkt.version

	for _, spkt := range pkt.subPackets {
		sum += sumVersionNumbers(spkt)
	}

	return sum
}

func readPacket(transmission string, bitsRead int, removeZeros bool) (Packet, int) {
	br := bitsRead
	var pkt Packet

	pkt.version, _ = strconv.ParseInt(transmission[br:br+3], 2, 64)
	br += 3
	pkt.typeID, _ = strconv.ParseInt(transmission[br:br+3], 2, 64)
	br += 3

	switch pkt.typeID {
	case 4:
		last := false

		for !last {
			leadingBit := transmission[br : br+1]
			if leadingBit == "0" {
				last = true
			}
			pkt.content = fmt.Sprintf("%s%s", pkt.content, transmission[br:br+5])
			br += 5
		}

		if removeZeros {
			br += 4 - ((br - bitsRead) % 4)
		}
	default:
		pkt.lengthType, _ = strconv.Atoi(transmission[br : br+1])
		br += 1
		switch pkt.lengthType {
		case 0:
			pkt.length, _ = strconv.ParseInt(transmission[br:br+15], 2, 64)
			br += 15
			pkt.content = transmission[br : br+int(pkt.length)]
			br += int(pkt.length)
			nbr := 0
			for nbr < len(pkt.content) {
				nPkt, nnbr := readPacket(pkt.content, nbr, false)
				pkt.subPackets = append(pkt.subPackets, nPkt)
				nbr = nnbr
			}
		case 1:
			pkt.length, _ = strconv.ParseInt(transmission[br:br+11], 2, 64)
			br += 11
			lbr := br
			for i := int64(0); i < pkt.length; i++ {
				nPkt, nbr := readPacket(transmission, br, false)
				pkt.subPackets = append(pkt.subPackets, nPkt)
				br = nbr
			}
			pkt.content = transmission[lbr:br]
		}

	}

	return pkt, br
}

func parseInputs(inp []string) string {
	byteHexMap := map[rune]string{
		'0': "0000", '1': "0001",
		'2': "0010", '3': "0011",
		'4': "0100", '5': "0101",
		'6': "0110", '7': "0111",
		'8': "1000", '9': "1001",
		'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101",
		'E': "1110", 'F': "1111",
	}
	res := ""
	for _, v := range inp[0] {
		res = fmt.Sprintf("%s%s", res, byteHexMap[v])
	}

	return res
}

func getInputsByLine() []string {
	inputFile, err := os.Open("./inputs.txt")
	if err != nil {
		fmt.Println("could not find file")
		os.Exit(1)
	}

	defer inputFile.Close()

	var inputs []string

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
