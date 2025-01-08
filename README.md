# OrGOnizer - File Organizer CLI

OrGOnizer is a simple command-line application written in Go to help you organize your files into categorized directories based on file extensions. It's a handy tool to tidy up directories with minimal effort.

## Features

- Organizes files into categories:
  - **Images**: `.jpg`, `.jpeg`, `.png`, `.avif`, `.webp`, `.gif`, `.svg`
  - **Videos**: `.mp4`, `.mov`, `.mkv`
  - **Design**: `.psd`, `.ai`, `.xcf`
  - **Documents**: `.txt`, `.csv`, `.docx`, `.doc`, `.xlsx`, `.xls`, `.pdf`, `.odt`, `.ods`
  - **Others**: Files with extensions not matching the above categories
- Automatically creates folders for each category and moves the files into them.
- Lightweight and easy to use.

## Prerequisites

- Go (1.18 or later) installed on your system.

## Installation

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Build the executable:

   ```bash
   go build -o OrGOnizer
   ```

3. Move the executable to a directory in your PATH (optional):

   ```bash
   mv OrGOnizer /usr/local/bin/
   ```

## Usage

1. Run the application with the directory path you want to organize:

   ```bash
   OrGOnizer [directory path]
   ```

   For example:

   ```bash
   OrGOnizer ~/Downloads
   ```

2. The application will categorize the files in the specified directory and move them into corresponding folders.

## Example Output

If the `~/Downloads` directory contains:

- `image1.jpg`
- `video1.mp4`
- `document1.pdf`
- `unknown.xyz`

After running the tool:

```
~/Downloads
├── Images
│   └── image1.jpg
├── Videos
│   └── video1.mp4
├── Documents
│   └── document1.pdf
└── Others
    └── unknown.xyz
```

## Error Handling

- If a directory cannot be read, an error message will be displayed.
- If a file cannot be moved, the application will print an error message and continue processing other files.

## Contribution

Contributions are welcome! Feel free to submit issues or pull requests to improve OrGOnizer.

## License

This project is licensed under the MIT License.

---

Start organizing your files effortlessly with OrGOnizer! 🎉