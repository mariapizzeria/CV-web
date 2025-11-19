import { EducationApiResponse } from './interfaces.js';
import { SkillApiResponse } from './interfaces.js';

export class Api {
    private baseUrl: string;

    constructor(baseUrl: string = 'http://localhost:8082') {
        this.baseUrl = baseUrl;
    }

    async getEducation(): Promise<EducationApiResponse[]> {
        try {
            const response = await fetch(`${this.baseUrl}/education`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error('Error fetching education api', error);
            throw error;
        }
    }

    async getSkills(): Promise<SkillApiResponse[]> {
        try {
            const response = await fetch(`${this.baseUrl}/skills`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error('Error fetching education api', error);
            throw error;
        }
    }
}