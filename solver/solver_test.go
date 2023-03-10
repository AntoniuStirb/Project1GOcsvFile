package solver

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDeleteInvalidLines(t *testing.T) {

	type deleteInvalidLinesTest struct {
		name           string
		input          [][]string
		expectedOutput [][]string
	}

	var insertedRecords = [][]string{
		{"id", "first_name", "last_name", "email", "gender", "ip_address"},
		{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
		{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "", "9.254.198.50"},
		{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		{"5", "Benoite", "Jaffray", "bjaffray4@github.com", "Female", ""},
		{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"},
		{"8", "Andrea", "", "amckeran7@example.com", "", ""},
		{"", "Marnia", "De Roberto", "mderoberto8@mac.com", "Female", "99.28.96.140"},
		{"10", "Marnia", "De Roberto", "mderoberto8@mac.com"},
		{},
	} //avem mai putine records, trebuie eliminat

	var expectedRecords = [][]string{
		{"id", "first_name", "last_name", "email", "gender", "ip_address"},
		{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
		{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"},
	}

	var deleteInvalidLineTests = []deleteInvalidLinesTest{
		{
			"when the input is empty, the output should be empty",
			nil,
			nil,
		},
		{
			"when the input is non-empty, the output should be not empty",
			insertedRecords,
			expectedRecords,
		},
	}

	for _, test := range deleteInvalidLineTests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := DeleteInvalidLines(test.input)
			if diff := cmp.Diff(actualOutput, test.expectedOutput); diff != "" {
				t.Errorf("TestedDeleteInvalidLines() does not meet expectations, "+
					"\nactual=%#v, \nexpected=%#v, \nDIFF: %v", actualOutput, test.expectedOutput, diff)
			}
		})
	}
}

func TestSplitFile(t *testing.T) {
	var splitRecordsTest1 = [][]string{
		{"id", "first_name", "last_name", "email", "gender", "ip_address"},
		{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
		{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "Male", "9.254.198.50"},
		{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		{"5", "Benoite", "Jaffray", "bjaffray4@github.com", "Female", "72.200.10.95"},
		{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"},
		{"8", "Andrea", "John", "amckeran7@example.com", "Female", "72.200.10.92"},
		{"9", "Marnia", "De Roberto", "mderoberto8@mac.com", "Female", "99.28.96.140"},
	}

	var expectedOutput1 = [][][]string{
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
			{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "Male", "9.254.198.50"},
			{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"5", "Benoite", "Jaffray", "bjaffray4@github.com", "Female", "72.200.10.95"},
			{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"},
			{"8", "Andrea", "John", "amckeran7@example.com", "Female", "72.200.10.92"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"9", "Marnia", "De Roberto", "mderoberto8@mac.com", "Female", "99.28.96.140"},
		},
	}

	var splitRecordsTest2 = [][]string{
		{"id", "first_name", "last_name", "email", "gender", "ip_address"},
		{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
		{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "Male", "9.254.198.50"},
	}

	var splitRecordsTest3 = [][]string{
		{},
	}

	type args struct {
		chunkSize int
		records   [][]string
	}
	tests := []struct {
		name string
		args args
		want [][][]string
		//wantErr bool
	}{
		{
			name: "test_1",
			args: args{2, splitRecordsTest1},
			want: expectedOutput1,
		},
		{
			name: "test_2 ChunkSize equal with length of records",
			args: args{3, splitRecordsTest2},
			want: nil,
		},
		{
			name: "test_3 The input records is empty",
			args: args{0, splitRecordsTest3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitFile(tt.args.chunkSize, tt.args.records)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitFile() \ngot = %v, \nwant = %v", got, tt.want)
			}
		})
	}
}

type fakeReader struct {
	r           io.Reader
	returnedErr error
}

func (f *fakeReader) Read(p []byte) (n int, err error) {
	if f.returnedErr != nil {
		return n, f.returnedErr
	}
	return f.r.Read(p)

}

func TestReadCsv(t *testing.T) {
	tests := []struct {
		name    string
		input   io.Reader
		want    [][]string
		wantErr bool
	}{
		{
			name:    "Error flow test",
			input:   &fakeReader{returnedErr: errors.New("Test error")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Flow without error",
			input:   &fakeReader{r: strings.NewReader("tata,are,pere")},
			want:    [][]string{{"tata", "are", "pere"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCsv(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCsv() got = %v, want %v", got, tt.want)
			}
		})
	}
}
