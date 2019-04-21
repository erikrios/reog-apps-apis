package com.erikriosetiawan.perasaanku;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    ImageView imagePerasaan;
    TextView textPerasaan;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        getSupportActionBar().setTitle("Perasaan Saya");
        getSupportActionBar().setSubtitle("Ini adalah perasaan saya");

    }

    public void buttonClick(View view) {

        imagePerasaan = findViewById(R.id.img_happy);
        textPerasaan = findViewById(R.id.text_happy);

        textPerasaan.setText(R.string.menangis);
        imagePerasaan.setImageResource(R.drawable.img_menangis);

    }

    public void buttonClickMore(View view) {

        textPerasaan.setText(R.string.tertawa);
        imagePerasaan.setImageResource(R.drawable.img_tertawa);
    }
}
