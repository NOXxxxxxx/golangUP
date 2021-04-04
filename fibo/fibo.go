package fibo

func Fbn(n int) []uint64 {
	fbnq := make([]uint64,n)
	fbnq[0] =1
	fbnq[1]=1

	for i:=2;i<n;i++{
		fbnq[i]=fbnq[i-1]+fbnq[i-2]
	}

	return fbnq
}