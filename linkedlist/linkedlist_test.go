package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.Append(tt.args.value)
		})
	}
}

func TestLinkedList_BubbleSort(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if err := l.BubbleSort(); (err != nil) != tt.wantErr {
				t.Errorf("BubbleSort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_DeleteNode(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.DeleteNode(tt.args.node)
		})
	}
}

func TestLinkedList_Find(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if got := l.Find(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		node  *Node
		value int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if err := l.InsertAfter(tt.args.node, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InsertAfter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Merge(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		l2 *LinkedList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if err := l1.Merge(tt.args.l2); (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_MergeSort(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if err := l.MergeSort(); (err != nil) != tt.wantErr {
				t.Errorf("MergeSort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.Prepend(tt.args.value)
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if err := l.Reverse(); (err != nil) != tt.wantErr {
				t.Errorf("Reverse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Values(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			if got := l.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_appendDouble(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.appendDouble(tt.args.value)
		})
	}
}

func TestLinkedList_appendSingle(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.appendSingle(tt.args.value)
		})
	}
}

func TestLinkedList_prependDouble(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.prependDouble(tt.args.value)
		})
	}
}

func TestLinkedList_prependSingle(t *testing.T) {
	type fields struct {
		head     *Node
		tail     *Node
		isDouble bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:     tt.fields.head,
				tail:     tt.fields.tail,
				isDouble: tt.fields.isDouble,
			}
			l.prependSingle(tt.args.value)
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	type args struct {
		isDouble bool
	}
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.isDouble); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
