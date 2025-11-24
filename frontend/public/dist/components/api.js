export class Api {
    constructor(baseUrl = 'http://localhost:8082') {
        this.baseUrl = baseUrl;
    }
    async getEducation() {
        try {
            const response = await fetch(`${this.baseUrl}/education`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        }
        catch (error) {
            console.error('Error fetching education api', error);
            throw error;
        }
    }
    async getSkills() {
        try {
            const response = await fetch(`${this.baseUrl}/skills`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        }
        catch (error) {
            console.error('Error fetching education api', error);
            throw error;
        }
    }
    async getExperience() {
        try {
            const response = await fetch(`${this.baseUrl}/experience`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        }
        catch (error) {
            console.error('Error fetching experience api', error);
            throw error;
        }
    }
    async getSocial() {
        try {
            const response = await fetch(`${this.baseUrl}/social`);
            if (!response.ok) {
                throw new Error(`Http error: ${response.status}`);
            }
            return await response.json();
        }
        catch (error) {
            console.error('Error fetching social api', error);
            throw error;
        }
    }
}
