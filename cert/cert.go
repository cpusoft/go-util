package cert

import (
	"crypto/x509"
	"encoding/asn1"
	"io/ioutil"

	belogs "github.com/astaxie/beego/logs"
)

// fatherCerFile is as root, childCerFile is to be verified
// result: ok/fail
func VerifyCertsByX509(fatherCertFile string, childCertFile string) (result string, err error) {
	belogs.Debug("VerifyCertsByX509():fatherCertFile:", fatherCertFile, "   childCertFile:", childCertFile)

	fatherPool := x509.NewCertPool()
	fatherFileByte, err := ioutil.ReadFile(fatherCertFile)
	if err != nil {
		belogs.Debug("VerifyCertsByX509():fatherCertFile:", fatherCertFile, "   ReadFile err:", err)
		return "fail", err
	}
	faterCert, err := x509.ParseCertificate(fatherFileByte)
	if err != nil {
		belogs.Debug("VerifyCertsByX509():fatherCertFile:", fatherCertFile, "   ParseCertificate err:", err)
		return "fail", err
	}
	faterCert.UnhandledCriticalExtensions = make([]asn1.ObjectIdentifier, 0)
	belogs.Debug("VerifyCertsByX509():father issuer:", faterCert.Issuer.String(), "   subject:", faterCert.Subject.String())
	fatherPool.AddCert(faterCert)

	childFileByte, err := ioutil.ReadFile(childCertFile)
	if err != nil {
		belogs.Debug("VerifyCertsByX509():childCertFile:", childCertFile, "   ReadFile err:", err)
		return "fail", err
	}
	childCert, err := x509.ParseCertificate(childFileByte)
	if err != nil {
		belogs.Debug("VerifyCertsByX509():childCertFile:", childCertFile, "   ParseCertificate err:", err)
		return "fail", err
	}
	childCert.UnhandledCriticalExtensions = make([]asn1.ObjectIdentifier, 0)
	belogs.Debug("VerifyCertsByX509():child issuer:", childCert.Issuer.String(), "   childCert:", childCert.Subject.String())

	opts := x509.VerifyOptions{
		Roots: fatherPool,
		//Intermediates: inter,
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		//KeyUsages: []x509.ExtKeyUsage{x509.KeyUsageCertSign},
	}
	if _, err := childCert.Verify(opts); err != nil {
		belogs.Debug("VerifyCertsByX509():Verify err:", err)
		return "fail", err
	}
	return "ok", nil
}
