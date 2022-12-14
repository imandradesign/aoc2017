package main

import (
	"fmt"
	"strconv"
)

const (
	input = "951484596541141557316984781494999179679767747627132447513171626424561779662873157761442952212296685573452311263445163233493199211387838461594635666699422982947782623317333683978438123261326863959719777179228599319321138948466562743761584836184512984131635354116264899181952748224523953976485816295227945792555726121913344959454458829485471174415775278865324142733339789878929596275998341778873889585819916457474773252249179366599951454182657225576277834669222982366884688565754691273745959468648957498511326215934353963981471593984617554514519623785326888374742147318993423214834751785956958395133486656388454552769722562524415715913869946325551396638593398729938526424994348267935153555851552287223313383583669912941364344694725478258297498969517632881187394141593479818536194597976519254215932257653777455227477617957833273463216593642394215275314734914719726618923177918342664351954252667253233858814365351722938716621544226598956257753212248859258351363174782742336961425325381561575992352415514168782816173861148859478285339529151631429536819286498721812323861771638574344416879476255929929157912984151742613268754779685396125954595318134933366626594498249956388771723777242772654678448815844555372892574747735672368299826548254744359377667294764559334659523233146587568261116253155189394188696831691284711264872914348961888253386971994431352474717376878745948769171243242621219912378731755544387249443997382399714738351857752329367997665166956467544459817582915478514486541453932175598413554259672117364863112592515988922747164842668361925135551248923449968328385889877512156952725198691746951431443497496455761516486573476185321748523644283494181119399874324683922393547682851931435931276267766772798261563117954648576421741384823494187895272582575669685279986988357796138794326125852772995446355723211161523161886222562853546488411563473998633847953246787557146187696947831335722888918172961256498971868946237299523474841983527391489962357196433927251798764362493965894995592683296651874787384247326643886774966828657393717626591578321174832222434128817871765347278152799425565633521152643686221411129463425496425385516719682884157452772141585743166647191938727971366274357874252166721759"
)

//--------------------------------- Part One ---------------------------------//
// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list. The
// list is circular, so the digit after the last digit is the first digit in the
// list.
//
// For example:
//
//     1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
//     1111 produces 4 because each digit (all 1) matches the next.
//     1234 produces 0 because no digit matches the next.
//     91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

func partOne(numbers string) int {
	theSum := 0
	numberOfNumbers := len(numbers)

	for idx := 0; idx < numberOfNumbers-1; idx++ {
		theRune := numbers[idx]
		theInt, err := strconv.Atoi(string(theRune))

		if err != nil {
			panic(fmt.Sprintf("Hit an error, boo! %+v", err))
		}

		if theRune == numbers[idx+1] {
			// fmt.Printf("We're casting rune %c to %d\n", theRune, theInt)
			theSum += theInt
		}
	}

	if numbers[numberOfNumbers-1] == numbers[0] {
		theFirstNumber, err := strconv.Atoi(string(numbers[0]))

		if err != nil {
			panic("SAD")
		}

		theSum += theFirstNumber
	}

	return theSum
}

// Now, instead of considering the next digit, it wants you to consider the
// digit halfway around the circular list. That is, if your list contains 10
// items, only include a digit in your sum if the digit 10/2 = 5 steps forward
// matches it. Fortunately, your list has an even number of elements.

// For example:

//     1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
//     1221 produces 0, because every comparison is between a 1 and a 2.
//     123425 produces 4, because both 2s match each other, but no other digit has a match.
//     123123 produces 12.
//     12131415 produces 4.

func addUpStuff(numbers string, offset int) int {
	theSum := 0
	numberOfNumbers := len(numbers)

	for idx := 0; idx < numberOfNumbers-1; idx++ {
		theRune := numbers[idx]
		theInt := beAnIntDamnit(theRune)

		// This is where it gets messy AF
		compareWithIndex := idx + offset
		if compareWithIndex > numberOfNumbers-1 {
			compareWithIndex = compareWithIndex - numberOfNumbers
		}

		if theRune == numbers[compareWithIndex] {
			theSum += theInt
		}
	}

	return theSum
}

func partTwo(numbers string) int {
	return addUpStuff(numbers, len(numbers)/2)
}

func beAnIntDamnit(charcter byte) int {
	theInt, err := strconv.Atoi(string(charcter))

	if err != nil {
		panic("Failed to parse an int")
	}

	return theInt
}

func main() {
	fmt.Println("It's the AOC!")

	fmt.Printf("Part One Result: %d\n", partOne(input))
	fmt.Printf("Part Two Result: %d\n", partTwo(input))
}
