interface Doctor {
	id: number;
	firstName: string;
	lastName: string;
	dni: string;
	email: string;
	licenseNumber: string;
}

export const Doctoresdata: Doctor[] = [
	{ id: 1, firstName: "Carlos", lastName: "Pérez", dni: "30.123.456", email: "carlos.perez@gmail.com", licenseNumber: "MN 12345" },
	{ id: 2, firstName: "María", lastName: "González", dni: "31.987.654", email: "maria.gonzalez@gmail.com", licenseNumber: "MN 23456" },
	{ id: 3, firstName: "Jorge", lastName: "Rodríguez", dni: "28.456.789", email: "jorge.rodriguez@gmail.com", licenseNumber: "MN 34567" },
	{ id: 4, firstName: "Lucía", lastName: "Fernández", dni: "33.741.852", email: "lucia.fernandez@gmail.com", licenseNumber: "MN 45678" },
	{ id: 5, firstName: "Martín", lastName: "López", dni: "29.159.753", email: "martin.lopez@gmail.com", licenseNumber: "MN 56789" },
	{ id: 6, firstName: "Sofía", lastName: "Martínez", dni: "32.963.852", email: "sofia.martinez@gmail.com", licenseNumber: "MN 67890" },
];