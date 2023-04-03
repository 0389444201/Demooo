package utils

import (
	"demo/userAccount/api/internal/types"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportFile(data []types.ResponseGetAll) (err error) {

	err = setValues(data)

	if err != nil {
		return err
	}
	return nil
}

func setValues(data []types.ResponseGetAll) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.SetCellValue("sheet1", "A1", "ID")
	f.SetCellValue("sheet1", "B1", "Name")
	f.SetCellValue("sheet1", "C1", "Gender")
	f.SetCellValue("sheet1", "D1", "Email")

	for i := 0; i < len(data); i++ {
		row1 := fmt.Sprintf("A%v", i+2)
		row2 := fmt.Sprintf("B%v", i+2)
		row3 := fmt.Sprintf("C%v", i+2)
		row4 := fmt.Sprintf("D%v", i+2)
		f.SetCellValue("sheet1", row1, data[i].Id)
		f.SetCellValue("sheet1", row2, data[i].Name)
		f.SetCellValue("sheet1", row3, data[i].Gender)
		f.SetCellValue("sheet1", row4, data[i].Email)
	}

	if err := f.SaveAs("data.xlsx"); err != nil {
		return err
	}
	return nil
}
