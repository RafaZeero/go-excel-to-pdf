package main

import (
	"encoding/xml"
	"os"
)

type Presentation struct {
	XMLName        xml.Name `xml:"p:presentation"`
	XmlnsA         string   `xml:"xmlns:a,attr"`
	XmlnsR         string   `xml:"xmlns:r,attr"`
	XmlnsP         string   `xml:"xmlns:p,attr"`
	XmlnsC         string   `xml:"xmlns:c,attr"`
	XmlnsW         string   `xml:"xmlns:w,attr"`
	XmlnsMso       string   `xml:"xmlns:mso,attr"`
	XmlnsVt        string   `xml:"xmlns:vt,attr"`
	XmlnsO         string   `xml:"xmlns:o,attr"`
	XmlnsXsi       string   `xml:"xmlns:xsi,attr"`
	XmlnsPkg       string   `xml:"xmlns:pkg,attr"`
	Ps             string   `xml:"ps,attr"`
	SViewPr        SViewPr
	SldIdLst       SldIdLst
	SldMasterIdLst SldMasterIdLst
	SldSz          SldSz
}

type SViewPr struct {
	Cont bool `xml:"cont,attr"`
}

type SldIdLst struct {
	SldId SldId
}

type SldId struct {
	Id string `xml:"id,attr"`
}

type SldMasterIdLst struct {
}

type SldSz struct {
	Cx   string `xml:"cx,attr"`
	Cy   string `xml:"cy,attr"`
	Type string `xml:"type,attr"`
}

func main() {
	// Create a new presentation structure.
	presentation := Presentation{
		XmlnsA:   "http://schemas.openxmlformats.org/drawingml/2006/main",
		XmlnsR:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		XmlnsP:   "http://schemas.openxmlformats.org/presentationml/2006/main",
		XmlnsC:   "http://schemas.openxmlformats.org/drawingml/2006/chart",
		XmlnsW:   "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		XmlnsMso: "urn:schemas-microsoft-com:office:office",
		XmlnsVt:  "urn:schemas-microsoft-com:office:publisching",
		XmlnsO:   "urn:schemas-microsoft-com:office:office",
		XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsPkg: "http://schemas.microsoft.com/office/2006/xmlPackage",
		Ps:       "96",
		SViewPr: SViewPr{
			Cont: true,
		},
		SldIdLst: SldIdLst{
			SldId: SldId{
				Id: "256",
			},
		},
		SldMasterIdLst: SldMasterIdLst{},
		SldSz: SldSz{
			Cx:   "12192000",
			Cy:   "6858000",
			Type: "screen16x9",
		},
	}

	// Create and open the PPTX file.
	file, err := os.Create("presentation.pptx")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Marshal the presentation structure into XML.
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(presentation); err != nil {
		panic(err)
	}
}
