package models

type rawFile struct {
    name string
    rawURL string
    data []byte
}

func (rf *rawFile) Name() string  {
    return rf.name
}

func (rf *rawFile) RawUrl() string  {
    return rf.rawURL
}

func (rf *rawFile) Data() []byte {
    return rf.data
}

func (rf *rawFile) SetData(p []byte) {
    rf.data=p
}