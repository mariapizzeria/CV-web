import {EducationApiResponse} from "./payload";

export class EducationApi {
    private baseUrl: string;
    constructor(baseUrl: string = 'http://localhost:8082') {
        this.baseUrl = baseUrl;

    }

    async getEducation(): Promise<EducationApiResponse> {
        console.log(1);
        try {
            const response = await fetch(`${this.baseUrl}/education`);

            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json()
        } catch (error) {
            console.error('Error fetching education api', error);
            throw error;
        }
    }
}