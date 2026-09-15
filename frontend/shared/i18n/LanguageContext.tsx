import { createContext, useContext, useEffect, useState } from "react";
import es from "./es.json";
import en from "./en.json";

export type Language = "es" | "en";

const translations: Record<Language, Record<string, unknown>> = { es, en };

const STORAGE_KEY = "hospital-lang";

interface LanguageContextType {
	lang: Language;
	setLanguage: (lang: Language) => void;
	t: (key: string) => string;
}

const LanguageContext = createContext<LanguageContextType | undefined>(undefined);

const getNested = (obj: Record<string, unknown>, key: string): unknown =>
	key.split(".").reduce<unknown>((acc, part) => {
		if (acc && typeof acc === "object" && part in acc) {
			return (acc as Record<string, unknown>)[part];
		}
		return undefined;
	}, obj);

export const LanguageProvider = ({ children }: { children: React.ReactNode }) => {
	const [lang, setLang] = useState<Language>("es");

	useEffect(() => {
		const saved = localStorage.getItem(STORAGE_KEY);
		if (saved === "es" || saved === "en") {
			setLang(saved);
		}
	}, []);

	const setLanguage = (next: Language) => {
		setLang(next);
		localStorage.setItem(STORAGE_KEY, next);
		document.documentElement.lang = next;
	};

	useEffect(() => {
		document.documentElement.lang = lang;
	}, [lang]);

	const t = (key: string): string => {
		const value = getNested(translations[lang], key);
		if (typeof value === "string") return value;
		const fallback = getNested(translations.es, key);
		return typeof fallback === "string" ? fallback : key;
	};

	return (
		<LanguageContext.Provider value={{ lang, setLanguage, t }}>
			{children}
		</LanguageContext.Provider>
	);
};

export const useLanguage = (): LanguageContextType => {
	const ctx = useContext(LanguageContext);
	if (!ctx) {
		throw new Error("useLanguage debe usarse dentro de LanguageProvider");
	}
	return ctx;
};