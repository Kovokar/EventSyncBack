package utils

import (
	"regexp"
)

// removeMask remove todos os caracteres não numéricos de uma string
func removeMask(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}

// -------------------------
// CNPJ
// -------------------------

func AddMaskCNPJ(cnpj string) string {
	cnpj = removeMask(cnpj)
	if len(cnpj) != 14 {
		return cnpj
	}
	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}

func RmMaskCNPJ(cnpj string) string {
	return removeMask(cnpj)
}

// -------------------------
// CPF
// -------------------------

func AddMaskCPF(cpf string) string {
	cpf = removeMask(cpf)
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

func RmMaskCPF(cpf string) string {
	return removeMask(cpf)
}

// -------------------------
// CEP
// -------------------------

func AddMaskCEP(cep string) string {
	cep = removeMask(cep)
	if len(cep) != 8 {
		return cep
	}
	return cep[:5] + "-" + cep[5:]
}

func RmMaskCEP(cep string) string {
	return removeMask(cep)
}

// -------------------------
// Telefone
// -------------------------

func AddMaskPhone(phone string) string {
	phone = removeMask(phone)
	switch len(phone) {
	case 10: // (XX) XXXX-XXXX
		return "(" + phone[:2] + ") " + phone[2:6] + "-" + phone[6:]
	case 11: // (XX) XXXXX-XXXX
		return "(" + phone[:2] + ") " + phone[2:7] + "-" + phone[7:]
	default:
		return phone
	}
}

func RmMaskPhone(phone string) string {
	return removeMask(phone)
}

// -------------------------
// Data (YYYYMMDD → DD/MM/YYYY)
// -------------------------

func AddMaskDate(date string) string {
	date = removeMask(date)
	if len(date) != 8 {
		return date
	}
	return date[6:] + "/" + date[4:6] + "/" + date[:4]
}

func RmMaskDate(date string) string {
	return removeMask(date)
}

// -------------------------
// Texto genérico
// -------------------------

func RmMaskGeneric(text string) string {
	re := regexp.MustCompile(`[.\-/()\s]`)
	return re.ReplaceAllString(text, "")
}
