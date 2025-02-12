use eframe::{egui, epi, NativeOptions};

fn main() {
    let options = NativeOptions::default();
    eframe::run_native(
        Box::new(MyApp::default()),
        options,
    );
}

struct MyApp {
    label: String,
}

impl Default for MyApp {
    fn default() -> Self {
        Self {
            label: "".to_owned(),
        }
    }
}

impl epi::App for MyApp {
    fn update(&mut self, ctx: &egui::CtxRef, _frame: &mut epi::Frame) {
        egui::CentralPanel::default().show(ctx, |ui| {
            ui.label(&self.label);
            if ui.button("Quit").clicked() {
                std::process::exit(0);
            }
        });
    }

    fn name(&self) -> &str {
        "FLNGB"
    }
}