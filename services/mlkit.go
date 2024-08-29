package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"regexp"
	"strings"
	"tl_mlkit/models"

	vision "cloud.google.com/go/vision/apiv1"
)

var mlkit *models.Mlkit
var TYPES = [2]string{"OLD", "NEW"}

func DetectText(w io.Writer, file multipart.File) (*models.Mlkit, error) {
	mlkit = new(models.Mlkit)
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	if len(annotations) == 0 {
		return nil, err
	} else {
		fmt.Fprintln(w, "Text:")
		mlkit.Description = annotations[0].Description

		MlKitResults := strings.Split(annotations[0].Description, ": ")
		if len(MlKitResults) < 5 {
			mlkit.Type = TYPES[1]
			MlKitResults = strings.Split(annotations[0].Description, "\n")
			// fmt.Print(len(MlKitResults))
			// fmt.Print(MlKitResults[len(MlKitResults)-2])

			// isMatchedNidCase1, err := regexp.MatchString(`^[0-9]`, MlKitResults[len(MlKitResults)-1])
			// if err != nil {
			// 	return nil, err
			// }
			// if isMatchedNidCase1 {
			// 	mlkit.Nid = MlKitResults[len(MlKitResults)-1]
			// }

			// isMatchedNidCase2, err := regexp.MatchString(`^[0-9]`, MlKitResults[len(MlKitResults)-2])
			// if err != nil {
			// 	return nil, err
			// }
			// if isMatchedNidCase2 {
			// 	mlkit.Nid = MlKitResults[len(MlKitResults)-2]
			// }

			// isMatchedNidCase3, err := regexp.MatchString(`^[0-9]`, MlKitResults[len(MlKitResults)-3])
			// if err != nil {
			// 	return nil, err
			// }
			// if isMatchedNidCase3 {
			// 	mlkit.Nid = MlKitResults[len(MlKitResults)-3]
			// }

			for i := 0; i < len(MlKitResults); i++ {
				SetterMlkitForNew(i, MlKitResults[i], MlKitResults)
			}

		} else {

			mlkit.Type = TYPES[0]
			MlKitResults = strings.Split(annotations[0].Description, "\n")
			for i := 0; i < len(MlKitResults); i++ {
				SetterMlkit(i, MlKitResults[i], MlKitResults)
			}

		}

	}

	return mlkit, nil
}

func SetterMlkit(index int, MlKitResult string, MlKitResults []string) {
	fmt.Printf("%d ----%s \n", index, MlKitResult)

	matched, err := regexp.MatchString(`Date of Birth`, MlKitResult)

	if err != nil {
		fmt.Print(err)
	}

	if matched {
		mlkit.DateOfBirth = strings.Split(MlKitResult, "Date of Birth: ")[1]
	} else {

	}

	matched_id, err := regexp.MatchString(`ID NO:`, MlKitResult)

	if err != nil {
		fmt.Print(err)
	}
	if matched_id {
		mlkit.Nid = strings.Split(MlKitResult, "ID NO: ")[1]
	}
	// switch index {
	// case 1:
	// 	mlkit.NameBng = strings.Split(MlKitResult, "\n")[0]
	// case 2:
	// 	mlkit.Name = strings.Split(MlKitResult, "\n")[0]
	// case 3:
	// 	mlkit.FahtherName = strings.Split(MlKitResult, "\n")[0]
	// case 4:
	// 	mlkit.MotherName = strings.Split(MlKitResult, "\n")[0]
	// case 5:
	// //	mlkit.DateOfBirth = strings.Split(MlKitResult, "\n")[0]
	// case 6:
	// 	//	mlkit.Nid = strings.Split(MlKitResult, "\n")[0]
	// }
}

func SetterMlkitForNew(index int, MlKitResult string, MlKitResults []string) {
	fmt.Printf("%d ----%s \n", index, MlKitResult)

	matched, err := regexp.MatchString(`Date of Birth`, MlKitResult)

	if err != nil {
		fmt.Print(err)
	}

	if matched {
		dataOfBirth := strings.Split(MlKitResult, "Date of Birth ")
		if len(dataOfBirth) >= 2 {
			mlkit.DateOfBirth = dataOfBirth[1]
		} else {
			isMatcheddataOfBirthCase1, err := regexp.MatchString(`^[0-9]`, MlKitResults[index+1])
			if err != nil {
				fmt.Print(err)
			}
			if isMatcheddataOfBirthCase1 {
				mlkit.DateOfBirth = MlKitResults[index+1]
			}

		}
		//	mlkit.DateOfBirth =
	}

	textMatched_NID, err := regexp.MatchString(`NID `, MlKitResult)
	if textMatched_NID {
		getMatchNewNid(err, index, MlKitResult, MlKitResults)
		return
	}

	textMatched_ID, err := regexp.MatchString(`ID `, MlKitResult)
	if textMatched_ID {
		getMatchNewNid(err, index, MlKitResult, MlKitResults)
		return
	}

	textMatched_No, err := regexp.MatchString(`No `, MlKitResult)
	if textMatched_No {
		getMatchNewNid(err, index, MlKitResult, MlKitResults)
		return
	}

}

func getMatchNewNid(err error, index int, MlKitResult string, MlKitResults []string) string {
	if err != nil {
		fmt.Print(err)
	}

	matchNidNo := strings.Split(MlKitResult, "No ")

	if len(matchNidNo) >= 2 {
		mlkit.Nid = matchNidNo[1]
		return getNidSpaceRemove()
	}

	Nid := strings.Split(MlKitResult, ". ")

	if len(Nid) >= 2 {
		mlkit.Nid = Nid[1]
		return getNidSpaceRemove()
	}

	textMathNe_ := strings.Split(MlKitResult, "Ne ")

	if len(textMathNe_) >= 2 {
		mlkit.Nid = Nid[1]
		return getNidSpaceRemove()
	}

	if index+1 < len(MlKitResults) {
		isMatchedNidCase1stIndex, err := regexp.MatchString(`^[0-9]`, MlKitResults[index+1])
		if err != nil {
			fmt.Print(err)
		}
		if isMatchedNidCase1stIndex {
			mlkit.Nid = MlKitResults[index+1]
			return getNidSpaceRemove()
		}
	}

	if index+2 < len(MlKitResults) {
		isMatchedNidCase2ndIndex, err := regexp.MatchString(`^[0-9]`, MlKitResults[index+2])
		if err != nil {
			fmt.Print(err)
		}
		if isMatchedNidCase2ndIndex {
			if len(MlKitResults[index+2]) == 12 {
				mlkit.Nid = MlKitResults[index+2]
				return getNidSpaceRemove()
			}

		}
	}

	//	mlkit.Nid = strings.Split(MlKitResult, ".")[1]
	return " "
}
func getNidSpaceRemove() string {
	NidArray := strings.Split(mlkit.Nid, " ")
	if len(NidArray) >= 3 {
		mlkit.Nid = NidArray[0] + NidArray[1] + NidArray[2]
	}
	return mlkit.Nid
}
