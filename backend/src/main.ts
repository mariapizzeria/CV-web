// Объединяем все в одном файле чтобы избежать проблем с импортами

interface EducationApiResponse {
    duration: string;
    college: string;
    course: string;
}

class EducationApi {
    private baseUrl: string;

    constructor(baseUrl: string = 'http://localhost:8082') {
        this.baseUrl = baseUrl;
    }

    async getEducation(): Promise<EducationApiResponse[]> {
        console.log('Fetching from:', `${this.baseUrl}/education`);
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
}

class Renderer {
    private container: HTMLElement;

    constructor(containerId: string = 'education-container') {
        const container = document.getElementById(containerId);
        if (!container) {
            throw new Error(`No container with id ${containerId}`);
        }
        this.container = container;
    }

    public createEducationItem(edu: EducationApiResponse[]) {
        let htmlString = '';

        edu.forEach((item, index) => {
            const isLeft = index % 2 === 0;

            if (isLeft) {
                htmlString += `
            <div class="timeline-item flex flex-col md:flex-row items-center relative">
                <div class="timeline-content w-full md:w-1/2 text-right pr-8 md:pr-12 mb-4 md:mb-0">
                    <h3 class="caveat text-2xl md:text-3xl text-accent">${item.duration}</h3>
                    <span class="text-gray-600">${item.college}</span>
                    <p class="font-bold text-xl md:text-2xl text-accent mt-2">${item.course}</p>
                </div>
                <div class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-12 h-12 md:w-16 md:h-16 z-10">
                    <div class="w-4 h-4 md:w-6 md:h-6 bg-accent rounded-full border-4 border-white"></div>
                </div>
                <div class="w-full md:w-1/2 pl-8 md:pl-12"></div>
            </div>
            `;
            } else {
                htmlString += `
            <div class="timeline-item flex flex-col md:flex-row items-center relative">
                <div class="w-full md:w-1/2 pr-8 md:pr-12 mb-4 md:mb-0"></div>
                <div class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-12 h-12 md:w-16 md:h-16 z-10">
                    <div class="w-4 h-4 md:w-6 md:h-6 bg-accent rounded-full border-4 border-white"></div>
                </div>
                <div class="timeline-content w-full md:w-1/2 text-left pl-8 md:pl-12">
                    <h3 class="caveat text-2xl md:text-3xl text-accent">${item.duration}</h3>
                    <span class="text-gray-600">${item.college}</span>
                    <p class="font-bold text-xl md:text-2xl text-accent mt-2">${item.course}</p>
                </div>
            </div>
            `;
            }
        });

        this.container.innerHTML = htmlString;
    }

    showLoading(): void {
        this.container.innerHTML = `
            <div class="flex justify-center items-center py-8">
                <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-accent"></div>
            </div>`;
    }
}

class App {
    private educationApi: EducationApi;
    private renderer: Renderer;

    constructor() {
        this.educationApi = new EducationApi();
        this.renderer = new Renderer('education-container');
    }

    async init(): Promise<void> {
        if (document.getElementById('education-container')) {
            await this.loadEducationData();
        }
    }

    private async loadEducationData(): Promise<void> {
        try {
            this.renderer.showLoading();
            const educationData = await this.educationApi.getEducation();
            this.renderer.createEducationItem(educationData);
        } catch(error) {
            console.error('Failed to load education data:', error);
            this.renderer.showLoading();
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const app = new App();
    app.init().catch(console.error);
});