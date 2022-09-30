package chainofresponsibility

import "testing"

func TestDemo(t *testing.T) {
	kimi := NewCommonManager("Kimi")
	pinkman := NewDirector("Pinkman")
	nori := NewGeneralManager("Nori")

	kimi.SetSuperior(pinkman)
	pinkman.SetSuperior(nori)

	request := Request{"請假", "吱吱請假", 1}
	kimi.RequestApplications(request)

	request2 := Request{"請假", "吱吱請假", 4}
	kimi.RequestApplications(request2)

	request3 := Request{"加薪", "吱吱跪求加薪", 400}
	kimi.RequestApplications(request3)

	request4 := Request{"加薪", "吱吱跪求加薪", 600}
	kimi.RequestApplications(request4)

	request5 := Request{"泡咖啡", "吱吱求喝咖啡", 1}
	kimi.RequestApplications(request5)
}
